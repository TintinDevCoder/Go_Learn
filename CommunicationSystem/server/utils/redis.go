package utils

import (
	"github.com/redis/go-redis/v9"
)

func InitPool(address string, maxIdle, maxActive int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     "",
		DB:           0,
		PoolSize:     maxActive,
		MinIdleConns: maxIdle,
	})
}
