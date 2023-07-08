package game

import (
	"context"

	"github.com/10n1s-backend/cmd/model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type GameHandler struct {
	db *gorm.DB
}

type Config struct {
	Enabled bool `config:"enabled"`
}

type Interface interface {
	Create(ctx context.Context, hostID int, latitude, longitude decimal.Decimal) error

	Get(ctx context.Context, id int) (*model.Game, error)
	GetByUserID(ctx context.Context, id int) (*model.Game, error)

	List(ctx context.Context) ([]model.Game, error)
	ListNGamesByLocation(ctx context.Context, n int, latitude, longitude decimal.Decimal) ([]model.Game, error)

	Update(ctx context.Context, game *model.Game) error

	Delete(ctx context.Context, game *model.Game) error
}

func NewGameHandler(ctx context.Context, cfg Config) (Interface, error) {
	if cfg.Enabled {
		return &GameHandler{}, nil
	} else {
		return &dummy{}, nil
	}
}
