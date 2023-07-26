package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/10n1s-backend/internal/room/model"
	"github.com/redis/go-redis/v9"
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
	keyRoom := getKeyFormatRoom(roomID)
	val, err := redis.LRange(ctx, keyRoom, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("lrange error: %w", err)
	}
	ret := []int{}
	for _, roomIDString := range val {
		roomIDInt, err := strconv.Atoi(roomIDString)
		if err != nil {
			return nil, fmt.Errorf("strconv error: %w", err)
		}
		ret = append(ret, roomIDInt)
	}
	return ret, nil
}

func (r *roomCache) SetSession(ctx context.Context, roomID, userID int, redis *redis.Client) (*model.Session, error) {
	keyRoom := getKeyFormatRoom(roomID)
	valueUser := getValueFormatUser(userID)
	val, err := redis.LPush(ctx, keyRoom, valueUser).Result()
	fmt.Println(val)
	if err != nil {
		return nil, fmt.Errorf("lpush error: %w", err)
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
