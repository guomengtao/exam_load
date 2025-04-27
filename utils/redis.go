package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gin-go-test/config"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetEnv("REDIS_ADDR", "127.0.0.1:6379"),
		Password: config.GetEnv("REDIS_PASSWORD", ""), // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("❌ Redis connection failed:", err)
	} else {
		fmt.Println("✅ Redis connected")
	}
}