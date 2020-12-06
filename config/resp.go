package config

import (
	"context"
	"fmt"
	"github.com/aasumitro/ginca/src/domain"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
)

var respConn *redis.Client

func (config AppConfig) SetupRESPConnection() {
	// get redis host:port
	host := fmt.Sprintf("%s:%s",
		viper.GetString(`REDIS_HOST`),
		viper.GetString(`REDIS_PORT`))
	// get redis password
	pwd := viper.GetString(`REDIS_PASSWORD`)
	// get redis database position
	db, _ := strconv.Atoi(viper.GetString(`REDIS_DB`))
	// make redis client with option and set connection for global use
	conn := redis.NewClient(&redis.Options {
		Addr:		host,
		Password:	pwd,
		DB:			db,
	})
	// set the resp connection for global usage
	setRESPConnection(conn)
}

func setRESPConnection(currentRESPConnection *redis.Client) {
	respConn = currentRESPConnection
}

func (config AppConfig) GetRESPConnection() *redis.Client {
	return respConn
}

func (config AppConfig) GetRESPStatus(ctx context.Context) string {
	if err := respConn.Ping(ctx).Err(); err != nil {
		return domain.RedisUnavailable.Error()
	}

	return domain.RedisAvailable
}