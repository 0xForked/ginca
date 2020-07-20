package config

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
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
	if viper.GetString(`SERVER_ENV`) == "production" {
		log.Println("Service RUN on RELEASE mode")
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println("Service RUN on DEBUG mode")
	}
}

func (config AppConfig) GetServerPort() string {
	return port
}

// InitAppConfig initialize the app configuration
func InitAppConfig() *AppConfig {
	return &AppConfig{}
}