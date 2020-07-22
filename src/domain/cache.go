package domain

type RedisCacheContract interface {
	Set(key string, value interface{})
	GetObject(key string) *map[string]interface{}
	GetArray(key string) *[]map[string]interface{}
	Delete(key string)
	Ping() string
}