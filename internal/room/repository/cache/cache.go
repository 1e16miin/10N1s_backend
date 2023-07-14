package cache

import (
	"context"

	"github.com/10n1s-backend/internal/room/model"
	"github.com/go-redis/redis/v8"
)

/*
session

	key   : roomID
	value : userID (list)
*/

type RoomRepositoryCache interface {
}
type roomCache struct{}

func NewRepositoryCache() RoomRepositoryCache {
	return &roomCache{}
}

func (r *roomCache) GetUsersIDByRoomID(ctx context.Context, roomID int, redis *redis.Client) ([]int, error) {
	return nil, nil
}

func (r *roomCache) SetSession(ctx context.Context, userID, roomID int, redis *redis.Client) (*model.Session, error) {
	return nil, nil
}

func (r *roomCache) DeleteSession(ctx context.Context, userID int, redis *redis.Client) error {
	return nil
}
