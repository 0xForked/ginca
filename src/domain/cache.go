package domain

import "context"

type RedisCacheContract interface {
	Set(ctx context.Context, key string, value interface{})
	Get(ctx context.Context, key string) *interface{}
	Delete(ctx context.Context, key string)
}