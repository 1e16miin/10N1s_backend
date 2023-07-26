package room

import (
	"context"
	"fmt"

	"github.com/10n1s-backend/internal/room/model"
	"github.com/10n1s-backend/internal/room/repository/cache"
	"github.com/10n1s-backend/internal/room/repository/database"
	"github.com/10n1s-backend/pkg/logger"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type RoomService interface {
	GetAllRooms(ctx context.Context) ([]model.Room, error)
	CreateRoom(ctx context.Context, hostID int, latitude, longitude string) (*model.Room, error)
}

type roomService struct {
	roomRepository      database.RoomRepository
	roomRepositoryCache cache.RoomRepositoryCache

	db     *gorm.DB
	cache  *redis.Client
	logger logger.Logger
}

func NewService(roomRepository database.RoomRepository, roomRepositoryCache cache.RoomRepositoryCache, db *gorm.DB, redis *redis.Client, logger logger.Logger) RoomService {
	return &roomService{roomRepository: roomRepository, roomRepositoryCache: roomRepositoryCache, db: db, cache: redis, logger: logger}
}

func (s *roomService) CreateRoom(ctx context.Context, hostID int, latitude, longitude string) (*model.Room, error) {
	lat, err := decimal.NewFromString(latitude)
	if err != nil {
		return nil, fmt.Errorf("cannot parse coordinates: %w", err)
	}
	long, err := decimal.NewFromString(longitude)
	if err != nil {
		return nil, fmt.Errorf("cannot parse coordinates: %w", err)
	}
	room, err := s.roomRepository.CreateRoom(&model.Room{HostID: hostID, Latitude: lat, Longitude: long}, s.db)
	if err != nil {
		return nil, fmt.Errorf("db error: %w", err)
	}
	return room, nil
}

func (s *roomService) GetAllRooms(ctx context.Context) ([]model.Room, error) {
	rooms, err := s.roomRepository.GetAllRooms(s.db)
	if err != nil {
		return nil, fmt.Errorf("db error: %w", err)
	}
	return rooms, nil
}

func (s *roomService) JoinRoom(ctx context.Context, roomID, userID int) error {
	_, err := s.roomRepositoryCache.SetSession(ctx, roomID, userID, s.cache)
	if err != nil {
		return fmt.Errorf("cache error: %w", err)
	}
	return nil
}
