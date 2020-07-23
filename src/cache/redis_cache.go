package cache

import (
	"context"
	"encoding/json"
	"github.com/aasumitro/gorest/logs"
	"github.com/aasumitro/gorest/src/domain"
	"github.com/go-redis/redis/v8"
	"time"
)

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

func (cache redisCache) Set(ctx context.Context, key string, value interface{}) {
	jsonMarshal, err := json.Marshal(value)
	if err != nil {
		logs.AppError.Println(err)
		return
	}

	cache.redisClient.Set(ctx, key, jsonMarshal, cache.expires)
}

func (cache redisCache) Get(ctx context.Context, key string) *interface{} {
	val, err := cache.redisClient.Get(ctx, key).Result()
	if err != nil {
		logs.AppError.Println(err)
		return nil
	}

	var data interface{}

	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		logs.AppError.Println(err)
		return nil
	}

	return &data
}

func (cache redisCache) Delete(ctx context.Context, key ...string)  {
	cache.redisClient.Del(ctx, key...)
}

func (cache redisCache) IsExist(ctx context.Context, key string) bool {
	_, err := cache.redisClient.Get(ctx, key).Result()
	return err != redis.Nil
}

func (cache redisCache) Expire(ctx context.Context, key string, expiration time.Duration) {
	if cache.IsExist(ctx, key) {
		cache.redisClient.Expire(ctx, key, expiration)
	}
}