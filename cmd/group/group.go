package group

import "context"

type GroupHandler struct{}
type Interface interface {
}

func NewGroupHandler(ctx context.Context) (Interface, error) {
	return &GroupHandler{}, nil
}
