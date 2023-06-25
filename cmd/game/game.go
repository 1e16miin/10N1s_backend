package game

import (
	"context"

	"github.com/10n1s-backend/cmd/model"
)

type GameHandler struct{}
type Interface interface {
	Create(ctx context.Context) error
	Get(ctx context.Context) (*model.Game, error) // admin 또는 내부적으로 이 게임에 대해 접근
	List(ctx context.Context) ([]model.Game, error)
	Update(ctx context.Context) error
	Delete(ctx context.Context, game *model.Game) error
}

func NewGameHandler(ctx context.Context) (Interface, error) {
	return &GameHandler{}, nil
}
