package game

import (
	"context"
)

type gameHandler struct{}

type Config struct{}

type Interface interface {
	Start(ctx context.Context) error
	End(ctx context.Context) error
	Get(ctx context.Context) error
	List(ctx context.Context) error
}

func NewGameHandler(ctx context.Context, cfg Config) *gameHandler {
	return &gameHandler{}
}

func (g *gameHandler) Start(ctx context.Context) error {

	return nil
}

func (g *gameHandler) End(ctx context.Context) error {
	return nil
}

func (g *gameHandler) Get(ctx context.Context) error {
	return nil
}

func (g *gameHandler) List(ctx context.Context) error {
	return nil
}

/*
create(ctx context.Context, hostID int, latitude, longitude decimal.Decimal) error

get(ctx context.Context, id int) (*model.Game, error)
getByUserID(ctx context.Context, id int) (*model.Game, error)

list(ctx context.Context) ([]model.Game, error)
listGamesByLocation(ctx context.Context, n int, latitude, longitude decimal.Decimal) ([]model.Game, error)

update(ctx context.Context, game *model.Game) error

delete(ctx context.Context, game *model.Game) error
*/
