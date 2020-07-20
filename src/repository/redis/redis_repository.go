package redis

//var ctx = context.Background()

//type redisCache struct {
//	redisClient			*redis.Client
//}

//func NewRedisCache(redis *redis.Client) domain.RedisRepository {
//	return &redisCache{redisClient: redis}
//}

//func (redis redisCache) GetCache(key string) {
	//data, err := redis.redisClient.Get(ctx, key).Result()
	//if data == "" || err != nil {
	//	return nil
	//}
	//
	//entity := &entity.Account{}
	//jsonUnmarshalError := json.Unmarshal([]byte(data), entity)
	//if entity.ID == "" || jsonUnmarshalError != nil {
	//	return nil
	//}
	//
	//return entity
//}

//func (redis redisCache) SetCache(key string, val string) {
	//repos
//}