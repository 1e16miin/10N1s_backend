package room

import (
	"context"
	"fmt"
	"time"

	"github.com/10n1s-backend/internal/room/model"
	"github.com/10n1s-backend/internal/room/repository/cache"
	"github.com/10n1s-backend/internal/room/repository/database"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type RoomService interface {
	GetAllRooms(ctx context.Context, db *gorm.DB) ([]model.Room, error)
	CreateRoom(ctx context.Context, hostID, latitude, longitude string, db *gorm.DB) (*model.Room, error)
}

type roomService struct {
	roomRepository      database.RoomRepository
	roomRepositoryCache cache.RoomRepositoryCache
}

func NewService(roomRepository database.RoomRepository, roomRepositoryCache cache.RoomRepositoryCache) RoomService {
	return &roomService{roomRepository: roomRepository, roomRepositoryCache: roomRepositoryCache}
}

func (s *roomService) CreateRoom(ctx context.Context, hostID, latitude, longitude string, db *gorm.DB) (*model.Room, error) {
	lat, err := decimal.NewFromString(latitude)
	if err != nil {
		return nil, fmt.Errorf("cannot parse coordinates: %w", err)
	}
	long, err := decimal.NewFromString(longitude)
	if err != nil {
		return nil, fmt.Errorf("cannot parse coordinates: %w", err)
	}

	room := &model.Room{HostID: hostID, Latitude: lat, Longitude: long, CreatedAt: time.Now()}
	return s.roomRepository.CreateRoom(ctx, room, db)
}

func (s *roomService) GetAllRooms(ctx context.Context, db *gorm.DB) ([]model.Room, error) {
	return s.roomRepository.GetAllRooms(ctx, db)
}
