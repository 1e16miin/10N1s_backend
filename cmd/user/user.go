package user

import "context"

type UserHandler struct{}
type Interface interface {
}

func NewUserHandler(ctx context.Context) (Interface, error) {
	return &UserHandler{}, nil
}
