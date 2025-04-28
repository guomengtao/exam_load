package utils

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
    "gin-go-test/config"
)

// RedisClient is the Redis client used to interact with the Redis database.
var RedisClient *redis.Client

// Ctx is the context used for Redis operations.
var Ctx = context.Background()

// InitRedis initializes the Redis client with configurations and tests the connection.
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