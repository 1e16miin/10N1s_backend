package group

import "context"

type GroupHandler struct{}
type Interface interface {
}

type Config struct {
	Enabled bool `config:"enabled"`
}

func NewGroupHandler(ctx context.Context, cfg Config) (Interface, error) {
	if cfg.Enabled {
		return &GroupHandler{}, nil
	} else {
		return &dummy{}, nil
	}
}
