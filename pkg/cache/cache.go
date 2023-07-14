package cache

import "github.com/go-redis/redis/v8"

type Config struct {
	Engine string `config:"engine"`
}

func NewCache(cfg Config) (*redis.Client, error) {
	return nil, nil
}
