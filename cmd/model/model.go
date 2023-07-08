package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Status int

const (
	Creating Status = 1 + iota
	Running
)

type Game struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	HostID     int `gorm:"unique"`
	GuestID1   int `gorm:"unique"`
	GuestID2   int `gorm:"unique"`
	GuestID3   int `gorm:"unique"`
	Location   Location
	LocationID int
	Type       Type
	TypeID     int
	StatusID   int
	Created    time.Time
	Deleted    time.Time
}

type Type struct {
	ID   int
	Name string `gorm:"unique"`
}

type Location struct {
	ID        int
	Name      string          `gorm:"unique"`
	Latitude  decimal.Decimal `gorm:"type:decimal(7,6);"`
	Longitude decimal.Decimal `gorm:"type:decimal(7,6);"`
}
