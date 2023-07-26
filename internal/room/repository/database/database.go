package database

import (
	"fmt"

	"github.com/10n1s-backend/internal/room/model"
	"gorm.io/gorm"
)

type RoomRepository interface {
	CreateRoom(room *model.Room, db *gorm.DB) (*model.Room, error)
	DeleteRoom(roomID int, db *gorm.DB) error
	UpdateRoom(room *model.Room, db *gorm.DB) (*model.Room, error)
	GetRoom(roomID int, db *gorm.DB) (*model.Room, error)
	GetRoomByHostID(hostID int, db *gorm.DB) (*model.Room, error)
	GetAllRooms(db *gorm.DB) ([]model.Room, error)
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

func (r *roomDB) CreateRoom(room *model.Room, db *gorm.DB) (*model.Room, error) {
	result := db.Create(room)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return room, nil
}

func (r *roomDB) DeleteRoom(roomID int, db *gorm.DB) error {
	result := db.Delete(&model.Room{}, roomID)
	if result.Error != nil {
		return fmt.Errorf("queryFailed: %w", result.Error)
	}
	return nil
}

func (r *roomDB) UpdateRoom(room *model.Room, db *gorm.DB) (*model.Room, error) {
	result := db.Save(room)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return room, nil
}

func (r *roomDB) GetRoom(roomID int, db *gorm.DB) (*model.Room, error) {
	room := &model.Room{}
	result := db.First(room, roomID)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return room, nil
}

func (r *roomDB) GetRoomByHostID(hostID int, db *gorm.DB) (*model.Room, error) {
	room := &model.Room{}
	result := db.Where("host_id =?", hostID).First(room)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return room, nil
}

func (r *roomDB) GetAllRooms(db *gorm.DB) ([]model.Room, error) {
	rooms := []model.Room{}
	result := db.Find(&rooms)
	if result.Error != nil {
		return nil, fmt.Errorf("queryFailed: %w", result.Error)
	}
	return rooms, nil
}
