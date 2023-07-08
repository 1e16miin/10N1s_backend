package group

import "context"

type groupHandler struct{}
type Interface interface {
}

type Config struct{}

func NewGroupHandler(ctx context.Context, cfg Config) *groupHandler {
	return &groupHandler{}
}
