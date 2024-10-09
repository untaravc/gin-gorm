package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Address of the Redis server
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Test the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err) // handle error appropriately
	}
	log.Println("Connected to REDIS")
	return rdb
}
