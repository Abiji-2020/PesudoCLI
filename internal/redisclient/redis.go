/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/

package redisclient

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedisClient(addr string) *RedisClient {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return &RedisClient{
		ctx:    ctx,
		client: client,
	}
}
