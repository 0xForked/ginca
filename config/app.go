package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"time"
)

var port string

// AppConfig represent the data-struct for configuration
type AppConfig struct {
	// another stuff , may be needed by configuration
}

func init() {
	// load and read config file
	setConfigFile()
	// set server environment
	setServerEnvironment()
	// set server port
	setServerPort()
}

func setConfigFile() {
	// find environment file
	viper.SetConfigFile(`.env`)
	// error handling for specific case
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(".env file not found!, please copy .env.example and paste as .env")
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}
}

func setServerPort() {
	port = viper.GetString(`SERVER_PORT`)
}

func setServerEnvironment()  {
	serviceVersion := viper.GetString(`SERVER_VERSION`)
	if viper.GetString(`SERVER_ENV`) == "production" {
		log.Println(fmt.Sprintf(
			"Service RUN on RELEASE mode, Service Version: %s",
			serviceVersion))
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println(fmt.Sprintf(
			"Service RUN on DEBUG mode, Service Version: %s",
			serviceVersion))
		gin.SetMode(gin.DebugMode)
	}
}

func (config AppConfig) GetServerPort() string {
	return port
}

func (config AppConfig) GetCacheTTL() time.Duration {
	cacheTTL := viper.GetInt(`CACHE_TTL`)
	cacheDuration := time.Duration(cacheTTL)
	return cacheDuration * time.Minute
}

func (config AppConfig) IsDevelopmentMode() bool {
	if viper.GetString(`SERVER_ENV`) == "production" {
		return false
	}

	return true
}

// InitAppConfig initialize the app configuration
func InitAppConfig() *AppConfig {
	return &AppConfig{}
}