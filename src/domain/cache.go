package domain

import (
	"context"
	"time"
)

type RedisCacheContract interface {
	Set(ctx context.Context, key string, value interface{})
	Get(ctx context.Context, key string) *interface{}
	Delete(ctx context.Context, key ...string)
	IsExist(ctx context.Context, key string) bool
	Expire(ctx context.Context, key string, expiration time.Duration)
}