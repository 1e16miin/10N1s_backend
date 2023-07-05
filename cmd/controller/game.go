package controller

import (
	"context"

	"github.com/shopspring/decimal"
)

func (h *ControllerHandler) CreateGame(ctx context.Context, hostID int, latitude, longitude decimal.Decimal) error {
	return h.game.Create(ctx, hostID, latitude, longitude)
}
