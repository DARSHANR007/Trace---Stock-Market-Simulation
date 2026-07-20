package market

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetRedisClient(client *redis.Client) {
	RedisClient = client
}

func cacheGet(key string, destination any) bool {
	if RedisClient == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	value, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	if err := json.Unmarshal([]byte(value), destination); err != nil {
		return false
	}

	return true
}

func cacheSet(key string, value any, ttl time.Duration) {
	if RedisClient == nil {
		return
	}

	encoded, err := json.Marshal(value)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	_ = RedisClient.Set(ctx, key, encoded, ttl).Err()
}
