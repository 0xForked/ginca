package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type mainHandler struct {}

func NewMainHandler(router *gin.Engine) {
	handler := &mainHandler{}
	router.GET("/", handler.home)
	router.GET("/health", handler.ping)
}

func (handler mainHandler) home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"status" : http.StatusText(http.StatusOK),
		"message": fmt.Sprintf("Welcome to %s", viper.GetString(`SERVER_NAME`)),
	})
}

func (handler mainHandler) ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"status" : http.StatusText(http.StatusOK),
		"message": map[string]string{
			"service" : "Service is running well",
			"mysql" : "",
			"redis": "",
		},
	})
}