package config

import (
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
	setConfig()
	// notify user about server environment is release mode or debug mode
	if viper.GetString(`SERVER_ENV`) == "debug" {
		log.Println("Service RUN on DEBUG mode")
	}
	// set server port
	setServerPort(viper.GetString(`SERVER_PORT`))
}

func setConfig() {
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

func setServerPort(Port string) {
	port = Port
}

func (config AppConfig) GetServerPort() string {
	return port
}

// SetupAppConfig initialize the app configuration
func SetupAppConfig() *AppConfig {
	return &AppConfig{}
}