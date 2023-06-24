package rank

import "context"

type RankHandler struct{}
type Interface interface {
}

func NewRankHandler(ctx context.Context) (Interface, error) {
	return &RankHandler{}, nil
}
