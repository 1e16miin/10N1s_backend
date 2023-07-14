package service

import (
	"context"
	"time"

	"github.com/10n1s-backend/internal/room"
	roomModel "github.com/10n1s-backend/internal/room/model"
	"github.com/10n1s-backend/pkg/logger"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Service struct {
	db     *gorm.DB
	redis  *redis.Client
	logger logger.Logger

	roomSVC room.RoomService
}

func NewService(db *gorm.DB, redis *redis.Client, roomSVC room.RoomService, logger logger.Logger) *Service {
	return &Service{db: db, redis: redis, roomSVC: roomSVC, logger: logger}
}

func (s *Service) GetRooms() ([]roomModel.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	/*
		만약 tx가 필요하고, user 에 접근이 필요하다면,
		tx := s.db.Begin()
		committed := false
		defer func() {
			if !committed {
				tx.rollback()
			}
		}
		rooms, err := s.roomSVC.GetRooms(ctx, tx)
		if err != nil {

		}
		users, err := s.userSVC.GetUsers(ctx, tx)
		if err != nil {

		}

		tx.Commit()
		committed = true

		return rooms, nil
	*/
	rooms, err := s.roomSVC.GetAllRooms(ctx, s.db)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return rooms, nil
}

func (s *Service) CreateRoom(hostID, latitude, longitude string) (*roomModel.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	room, err := s.roomSVC.CreateRoom(ctx, hostID, latitude, longitude, s.db)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return room, nil
}
