package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func CreateRedisConnection() (*redis.Client, error) {
	ctx := context.Background()
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := db.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	return db, nil
}
