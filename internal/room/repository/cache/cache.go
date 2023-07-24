package cache

import (
	"context"
	"fmt"

	"github.com/10n1s-backend/internal/room/model"
	"github.com/go-redis/redis/v8"
)

/*
session

	key   : room/{roomID}
	value : user/{userID} (list)
*/

type RoomRepositoryCache interface {
	GetUsersIDByRoomID(ctx context.Context, roomID int, redis *redis.Client) ([]int, error)
	SetSession(ctx context.Context, roomID, userID int, redis *redis.Client) (*model.Session, error)
	DeleteSessionUserID(ctx context.Context, userID int, redis *redis.Client) error
}
type roomCache struct{}

func NewRepositoryCache() RoomRepositoryCache {
	return &roomCache{}
}

func (r *roomCache) GetUsersIDByRoomID(ctx context.Context, roomID int, redis *redis.Client) ([]int, error) {
	return nil, nil
}

func (r *roomCache) SetSession(ctx context.Context, roomID, userID int, redis *redis.Client) (*model.Session, error) {
	keyRoom := getKeyFormatRoom(roomID)
	valueUser := getValueFormatUser(userID)
	val, err := redis.LPush(ctx, keyRoom, valueUser).Result()
	fmt.Println(val)
	if err != nil {
		return nil, fmt.Errorf("l push error: %w", err)
	}

	return &model.Session{UserID: userID, RoomID: roomID}, nil
}

func (r *roomCache) DeleteSessionUserID(ctx context.Context, userID int, redis *redis.Client) error {
	return nil
}

func getKeyFormatRoom(roomID int) string {
	return fmt.Sprintf("room/%d", roomID)
}

func getValueFormatUser(userID int) string {
	return fmt.Sprintf("user/%d", userID)
}
