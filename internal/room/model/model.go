package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Room struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	HostID    string
	Latitude  decimal.Decimal `gorm:"type:decimal(7,6);"`
	Longitude decimal.Decimal `gorm:"type:decimal(7,6);"`
	CreatedAt time.Time
	DeletedAt time.Time
}
type Session struct {
	UserID int
	RoomID int
}
