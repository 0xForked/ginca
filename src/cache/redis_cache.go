package cache

import (
	"context"
	"encoding/json"
	"github.com/aasumitro/gorest/src/domain"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type redisCache struct {
	redisClient			*redis.Client
	expires				time.Duration
}

func NewRedisCache(
	redis *redis.Client,
	exp time.Duration,
) domain.RedisCacheContract {
	return &redisCache{redisClient: redis, expires: exp}
}

func (cache redisCache) Set(key string, value interface{}) {
	jsonMarshal, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	cache.redisClient.Set(ctx, key, jsonMarshal, cache.expires)
}

func (cache redisCache) GetObject(key string) *map[string]interface{} {
	val, err := cache.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	var example map[string]interface{}

	err = json.Unmarshal([]byte(val), &example)
	if err != nil {
		panic(err)
	}

	return &example
}

func (cache redisCache) GetArray(key string) *[]map[string]interface{} {
	val, err := cache.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	var examples []map[string]interface{}

	err = json.Unmarshal([]byte(val), &examples)
	if err != nil {
		panic(err)
	}

	return &examples
}

func (cache redisCache) Ping() string {
	if err := cache.redisClient.Ping(ctx).Err(); err != nil {
		return domain.RedisUnavailable.Error()
	}

	return domain.RedisAvailable
}