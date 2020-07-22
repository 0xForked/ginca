package handler

import (
	"fmt"
	"github.com/aasumitro/gorest/src/domain"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type mainHandler struct {
	redis		domain.RedisRepository
	mysqlStatus	string
}

func NewMainHandler(router *gin.Engine, redis domain.RedisRepository, mysqlStatus string) {
	handler := &mainHandler{redis: redis, mysqlStatus: mysqlStatus}
	router.GET("/", handler.home)
	router.GET("/health", handler.ping)
	router.NoRoute(handler.notFound)
}

func (handler mainHandler) home(context *gin.Context) {
	context.JSON(http.StatusOK, domain.Respond{
		Code : http.StatusOK,
		Status : http.StatusText(http.StatusOK),
		Data: fmt.Sprintf("Welcome to %s", viper.GetString(`SERVER_NAME`)),
	})
}

func (handler mainHandler) ping(context *gin.Context) {
	context.JSON(http.StatusOK, domain.Respond{
		Code: http.StatusOK,
		Status : http.StatusText(http.StatusOK),
		Data: map[string]string{
			"app" : domain.ServiceAvailable,
			"storage" : handler.mysqlStatus,
			"cache": handler.redis.Ping(),
		},
	})
}

func (handler mainHandler) notFound(context *gin.Context) {
	context.JSON(http.StatusNotFound, domain.Respond{
		Code: http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data: domain.RouteNotFound.Error(),
	})
}