package domain

type RedisRepository interface {
	Set(key string, value interface{})
	GetObject(key string) *map[string]interface{}
	GetArray(key string) *[]map[string]interface{}
	Ping() string
}