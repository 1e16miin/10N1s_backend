package user

import "context"

type UserHandler struct{}
type Interface interface {
}

type Config struct {
	Enabled bool `config:"enabled"`
}

func NewUserHandler(ctx context.Context, cfg Config) (Interface, error) {
	if cfg.Enabled {
		return &UserHandler{}, nil
	} else {
		return &dummy{}, nil
	}
}
