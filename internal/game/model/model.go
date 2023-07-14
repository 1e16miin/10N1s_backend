package model

import "time"

type Game struct {
	ID        int
	RoomID    int
	CreatedAt time.Time
	DeletedAt time.Time
}

type GameHistory struct { // mysql bulk create
	ID     int
	UserID int
	GameID int
	Result string
}
