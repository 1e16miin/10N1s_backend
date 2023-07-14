package database

import (
	"context"
	"fmt"

	"github.com/10n1s-backend/internal/room/model"
	"gorm.io/gorm"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *model.Room, db *gorm.DB) (*model.Room, error)
	DeleteRoom(ctx context.Context, roomID int, db *gorm.DB) error
	UpdateRoom(ctx context.Context, room *model.Room, db *gorm.DB) (*model.Room, error)
	GetRoom(ctx context.Context, roomID int, db *gorm.DB) (*model.Room, error)
	GetRoomByHostID(ctx context.Context, hostID int, db *gorm.DB) (*model.Room, error)
	GetAllRooms(ctx context.Context, db *gorm.DB) ([]model.Room, error)
}

type Config struct {
	DBInit bool `config:"dbInit"`
}

type roomDB struct{}

func NewRepository(cfg Config, db *gorm.DB) (RoomRepository, error) {
	r := &roomDB{}
	if cfg.DBInit {
		err := r.autoMigrateRoom(db)
		if err != nil {
			return nil, fmt.Errorf("cannot access to room table: %w", err)
		}
	}

	return r, nil
}

func (r *roomDB) autoMigrateRoom(db *gorm.DB) error {
	if db.Migrator().HasTable(&model.Room{}) {
		return nil
	}
	return db.Migrator().CreateTable(&model.Room{})
}

func (r *roomDB) CreateRoom(ctx context.Context, room *model.Room, db *gorm.DB) (*model.Room, error) {
	result := db.Create(room)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return room, nil
}

func (r *roomDB) DeleteRoom(ctx context.Context, roomID int, db *gorm.DB) error {
	result := db.Delete(&model.Room{}, roomID)
	if result.Error != nil {
		return fmt.Errorf("queryFailed: %w", result.Error)
	}
	return nil
}

func (r *roomDB) UpdateRoom(ctx context.Context, room *model.Room, db *gorm.DB) (*model.Room, error) {
	result := db.Save(room)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return room, nil
}

func (r *roomDB) GetRoom(ctx context.Context, roomID int, db *gorm.DB) (*model.Room, error) {
	room := &model.Room{}
	result := db.First(room, roomID)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return room, nil
}

func (r *roomDB) GetRoomByHostID(ctx context.Context, hostID int, db *gorm.DB) (*model.Room, error) {
	room := &model.Room{}
	result := db.Where("host_id =?", hostID).First(room)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return room, nil
}

func (r *roomDB) GetAllRooms(ctx context.Context, db *gorm.DB) ([]model.Room, error) {
	rooms := []model.Room{}
	result := db.Find(&rooms)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return rooms, nil
}
