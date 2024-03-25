package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"task/config"
)

func InitRedis(c *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       c.Redis.DBName,
	})
	_, err := rdb.Ping(context.Background()).Result()

	return rdb, err
}
