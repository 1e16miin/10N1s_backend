package controller

import (
	"context"

	"github.com/10n1s-backend/cmd/auth"
	"github.com/10n1s-backend/cmd/game"
	"github.com/10n1s-backend/cmd/group"
	"github.com/10n1s-backend/cmd/rank"
	"github.com/10n1s-backend/cmd/user"
)

type ControllerHandler struct {
	auth  auth.Interface
	game  game.Interface
	group group.Interface
	rank  rank.Interface
	user  user.Interface
}
type Interface interface {
}

func NewControllerHandler(ctx context.Context, auth auth.Interface, game game.Interface, group group.Interface, rank rank.Interface, user user.Interface) (Interface, error) {
	return &ControllerHandler{auth: auth, game: game, group: group, rank: rank, user: user}, nil
}
