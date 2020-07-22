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

func (cache redisCache) Get(key string) *interface{} {
	val, err := cache.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	var data interface{}

	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		panic(err)
	}

	return &data
}

func (cache redisCache) Delete(key string)  {
	cache.redisClient.Del(ctx, key)
}

func (cache redisCache) Ping() string {
	if err := cache.redisClient.Ping(ctx).Err(); err != nil {
		return domain.RedisUnavailable.Error()
	}

	return domain.RedisAvailable
}