/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/

package redisclient

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
)

func InitRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis:", err)
	}
	log.Printf("Connected to Redis at %s", "localhost:6379")
}
