package domain

type RedisCacheContract interface {
	Set(key string, value interface{})
	Get(key string) *interface{}
	Delete(key string)
	Ping() string
}