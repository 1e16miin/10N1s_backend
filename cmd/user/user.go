package user

import "context"

type userHandler struct{}
type Interface interface{}

type Config struct{}

func NewUserHandler(ctx context.Context, cfg Config) *userHandler {
	return &userHandler{}
}
