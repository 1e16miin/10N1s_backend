package rank

import "context"

type rankHandler struct{}
type Interface interface{}

type Config struct{}

func NewRankHandler(ctx context.Context, cfg Config) *rankHandler {
	return &rankHandler{}
}
