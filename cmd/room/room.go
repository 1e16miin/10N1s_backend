package room

import (
	"context"

	"github.com/10n1s-backend/cmd/model"
	"github.com/10n1s-backend/cmd/user"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type roomHandler struct {
	user user.Interface
	db   *gorm.DB
}

type Config struct{}

type Interface interface {
	Create(ctx context.Context, hostID int, latitude, longitude decimal.Decimal, db *gorm.DB) (int, error)

	Get(ctx context.Context, roomID int) error
	ListAll(ctx context.Context) ([]model.Room, error)
	ListNear(ctx context.Context, latitude, longitude decimal.Decimal) ([]model.Room, error)

	Enter(ctx context.Context, roomID, userID int) error
	Leave(ctx context.Context, userID int) error

	Delete(ctx context.Context, roomID int) error
}

func NewRoomHandler(ctx context.Context, cfg Config) *roomHandler {

	return &roomHandler{}
}

func (r *roomHandler) Create(ctx context.Context, hostID int, latitude, longitude decimal.Decimal, tx *gorm.DB) (int, error) {

	return 0, nil
}

func (r *roomHandler) ListAll(ctx context.Context, tx *gorm.DB) ([]model.Room, error) {
	// query
	return nil, nil
}

func (r *roomHandler) ListNear(ctx context.Context, latitude, longitude decimal.Decimal) ([]model.Room, error) {
	return nil, nil
}

func (r *roomHandler) Enter(ctx context.Context, roomID, userID int) error {

	return nil
}

func (r *roomHandler) Leave(ctx context.Context, userID int) error {
	return nil
}

func (r *roomHandler) Delete(ctx context.Context, roomID int) error {
	return nil
}
