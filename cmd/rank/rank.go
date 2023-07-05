package rank

import "context"

type RankHandler struct{}
type Interface interface {
}

type Config struct {
	Enabled bool `config:"enabled"`
}

func NewRankHandler(ctx context.Context, cfg Config) (Interface, error) {
	if cfg.Enabled {
		return &RankHandler{}, nil
	} else {
		return &dummy{}, nil
	}
}
