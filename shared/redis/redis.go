package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Ping(context.Background()).Err()
	return rdb, err
}
