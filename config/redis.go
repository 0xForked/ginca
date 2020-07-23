package config

import (
	"context"
	"fmt"
	"github.com/aasumitro/ginca/src/domain"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
)

var redisClient *redis.Client

func (config AppConfig) SetupRedisClientConnection() {
	// get redis host:port
	host := fmt.Sprintf("%s:%s",
		viper.GetString(`REDIS_HOST`),
		viper.GetString(`REDIS_PORT`))
	// get redis password
	pwd := viper.GetString(`REDIS_PASSWORD`)
	// get redis database position
	db, _ := strconv.Atoi(viper.GetString(`REDIS_DB`))
	// make redis client with option and set connection for global use
	redisClient = redis.NewClient(&redis.Options {
		Addr:		host,
		Password:	pwd,
		DB:			db,
	})
}

func (config AppConfig) GetRedisClientConnection() *redis.Client {
	return redisClient
}

func (config AppConfig) GetRedisStatus(ctx context.Context) string {
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return domain.RedisUnavailable.Error()
	}

	return domain.RedisAvailable
}