package game

import (
	"context"

	"github.com/10n1s-backend/cmd/model"
	"github.com/shopspring/decimal"
)

type dummy struct{}

func (d *dummy) Create(ctx context.Context, hostID int, latitude, longitude decimal.Decimal) error {
	return nil
}

func (d *dummy) Get(ctx context.Context, id int) (*model.Game, error) {
	return &model.Game{ID: 1}, nil
}

func (d *dummy) GetByUserID(ctx context.Context, id int) (*model.Game, error) {
	return &model.Game{ID: 1}, nil
}

func (d *dummy) List(ctx context.Context) ([]model.Game, error) {
	return nil, nil
}

func (d *dummy) ListNGamesByLocation(ctx context.Context, n int, latitude, longitude decimal.Decimal) ([]model.Game, error) {
	return nil, nil
}

func (d *dummy) Update(ctx context.Context, game *model.Game) error {
	return nil
}

func (d *dummy) Delete(ctx context.Context, game *model.Game) error {
	return nil
}
