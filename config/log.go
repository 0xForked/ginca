package config

import (
	"fmt"
	"github.com/aasumitro/ginca/logs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"time"
)

func (config AppConfig) SetupAppLog() {
	if config.IsDevelopmentMode() {
		localAccessLog()
		localErrorLog()
	}
}

func localAccessLog() {
	//Disable Console Color.
	gin.DisableConsoleColor()
	//Logging to a file.
	// file, _ := os.OpenFile(getLogName("access"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	file, err := os.Create(getLogName("access"))
	if err != nil {
		panic(err)
	}
	//write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func localErrorLog() {
	// file, _ := os.OpenFile(getLogName("errors"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	file, err := os.Create(getLogName("errors"))
	if err != nil {
		panic(err)
	}
	logs.AppError = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	logs.AppError.SetOutput(file)
	logs.AppError.Println("Initialize Application Error")
}

func sentryErrorLog() {
	// setup sentry error report
}

func getLogName(logType string) string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("./logs/%s/%v-%d-%d-%d.log",
		logType,
		viper.GetString(`SERVER_NAME`),
		year, month, day)
}