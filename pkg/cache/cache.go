package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Engine string      `config:"engine"`
	Redis  RedisConfig `config:"redis"`
}

type RedisConfig struct {
	Addr         string `config:"redisAddr"`
	Password     string `config:"password"`
	MinIdleConns int    `config:"minIdleConns"`
	PoolSize     int    `config:"poolSize"`
	PoolTimeout  int    `config:"poolTimeout"`
	DB           int    `config:"db"`
}

func NewCache(cfg Config) (*redis.Client, error) {
	redisHost := cfg.Redis.Addr
	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password, // no password set
		DB:           cfg.Redis.DB,       // use default DB
	})

	return client, nil
}
