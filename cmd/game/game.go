package game

import "context"

type GameHandler struct{}
type Interface interface {
}

func NewGameHandler(ctx context.Context) (Interface, error) {
	return &GameHandler{}, nil
}
