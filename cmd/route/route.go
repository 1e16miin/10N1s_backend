package route

import (
	"context"

	"github.com/10n1s-backend/cmd/controller"
)

type (
	Router interface {
		AddHandler(ctx context.Context) error
		Start() error
	}
	Config struct {
		Type string     `config:"type"`
		Echo EchoConfig `config:"echo"`
	}
)

func NewRouter(ctx context.Context, config Config, controller controller.Interface) (Router, error) {
	switch config.Type {
	case "echo":
		return NewEchoRouter(ctx, config.Echo, controller)
	default:
		return NewEchoRouter(ctx, config.Echo, controller)
	}
}
