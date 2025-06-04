package utils

import (
    "context"
    "fmt"
    "time"
    "github.com/redis/go-redis/v9"
    "gin-go-test/config"
)

// RedisClient is the Redis client used to interact with the Redis database.
var RedisClient *redis.Client

// Ctx is the context used for Redis operations.
var Ctx = context.Background()

// RedisHelper defines the interface for Redis operations
type RedisHelper interface {
	Get(ctx context.Context, key string) (string, error)
	Scan(ctx context.Context, matchPattern string) ([]string, error)
}

// DefaultRedisHelper implements RedisHelper using Redis client
type DefaultRedisHelper struct{}

func (h *DefaultRedisHelper) Get(ctx context.Context, key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

func (h *DefaultRedisHelper) Scan(ctx context.Context, matchPattern string) ([]string, error) {
	var keys []string
	var cursor uint64
	for {
		scannedKeys, nextCursor, err := RedisClient.Scan(ctx, cursor, matchPattern, 100).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, scannedKeys...)
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return keys, nil
}

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

// RedisSet sets a string value in Redis with expiration
func RedisSet(key string, value string, expiration time.Duration) error {
    return RedisClient.Set(Ctx, key, value, expiration).Err()
}

// RedisGet gets a string value from Redis
func RedisGet(key string) (string, error) {
    return RedisClient.Get(Ctx, key).Result()
}

// RedisDelete deletes a key from Redis
func RedisDelete(key string) error {
    return RedisClient.Del(Ctx, key).Err()
}

// RedisExists checks if a key exists in Redis
func RedisExists(key string) (bool, error) {
    result, err := RedisClient.Exists(Ctx, key).Result()
    return result == 1, err
}

// RedisHSet sets a hash field in Redis
func RedisHSet(key string, field string, value interface{}) error {
    return RedisClient.HSet(Ctx, key, field, value).Err()
}

// RedisHGet gets a hash field from Redis
func RedisHGet(key string, field string) (string, error) {
    return RedisClient.HGet(Ctx, key, field).Result()
}

// RedisHGetAll gets all hash fields from Redis
func RedisHGetAll(key string) (map[string]string, error) {
    return RedisClient.HGetAll(Ctx, key).Result()
}

// RedisHMSet sets multiple hash fields in Redis
func RedisHMSet(key string, fields map[string]interface{}) error {
    return RedisClient.HMSet(Ctx, key, fields).Err()
}

// RedisExpire sets the expiration time for a key
func RedisExpire(key string, expiration time.Duration) error {
    return RedisClient.Expire(Ctx, key, expiration).Err()
}