package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
	"time"
)

func (config AppConfig) SetupAppLog() {
	if config.IsDevelopmentMode() {
		localAccessLog()
	}
}

func localAccessLog() {
	//Disable Console Color.
	gin.DisableConsoleColor()
	//Logging to a file.
	file, _ := os.Create(getLogName())
	//write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func sentryErrorLog() {
	// setup sentry error report
}

func getLogName() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("./logs/%v-%d-%d-%d.log",
		viper.GetString(`SERVER_NAME`),
		year, month, day)
}
