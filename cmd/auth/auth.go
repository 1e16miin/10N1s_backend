package auth

import (
	"context"

	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}
type Interface interface {
}

type Config struct {
	Enabled bool `config:"enabled"`
}

func NewAuthHandler(ctx context.Context, cfg Config, db *gorm.DB) (Interface, error) {
	if cfg.Enabled {
		return &AuthHandler{db: db}, nil
	} else {
		return &dummy{}, nil
	}
}
