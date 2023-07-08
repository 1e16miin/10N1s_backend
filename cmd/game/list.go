package game

import (
	"context"

	"github.com/10n1s-backend/cmd/model"
	"github.com/shopspring/decimal"
)

func (g *GameHandler) List(ctx context.Context) ([]model.Game, error) {

	return nil, nil
}

func (g *GameHandler) ListNGamesByLocation(ctx context.Context, n int, latitude, longitude decimal.Decimal) ([]model.Game, error) {

	return nil, nil
}
