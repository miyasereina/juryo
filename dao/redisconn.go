package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func init() {
	addr := "127.0.0.1:6379"
	db := 0

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		DB:           db,
		PoolSize:     30,
		MinIdleConns: 10,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func RedisConn() *redis.Client {
	return RedisClient
}
