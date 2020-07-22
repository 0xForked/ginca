package handler

import (
	"fmt"
	"github.com/aasumitro/gorest/config"
	httpDelivery "github.com/aasumitro/gorest/src/delivery/http"
	"github.com/aasumitro/gorest/src/domain"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type mainHandler struct {
	config       *config.AppConfig
}

func NewMainHandler(
	router *gin.Engine,
	appConfig *config.AppConfig,
) {
	handler := &mainHandler{config: appConfig}
	router.GET("/", handler.home)
	router.GET("/health", handler.ping)
	router.NoRoute(handler.notFound)
}

func (handler mainHandler) home(context *gin.Context) {
	context.JSON(http.StatusOK, httpDelivery.Respond{
		Code : http.StatusOK,
		Status : http.StatusText(http.StatusOK),
		Data: fmt.Sprintf("Welcome to %s", viper.GetString(`SERVER_NAME`)),
	})
}

func (handler mainHandler) ping(context *gin.Context) {
	context.JSON(http.StatusOK, httpDelivery.Respond{
		Code: http.StatusOK,
		Status : http.StatusText(http.StatusOK),
		Data: map[string]string{
			"app" : domain.ServiceAvailable,
			"storage" : handler.config.GetDatabaseStatus(),
			"cache": handler.config.GetRedisStatus(context),
		},
	})
}

func (handler mainHandler) notFound(context *gin.Context) {
	context.JSON(http.StatusNotFound, httpDelivery.Respond{
		Code: http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data: domain.RouteNotFound.Error(),
	})
}