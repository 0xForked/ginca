package main

import (
	"github.com/aasumitro/gorest/config"
	dataCache "github.com/aasumitro/gorest/src/cache"
	httpHandler "github.com/aasumitro/gorest/src/delivery/http/handler"
	"github.com/aasumitro/gorest/src/delivery/http/middleware"
	dataSourceMySQL "github.com/aasumitro/gorest/src/repository/mysql"
	useCase "github.com/aasumitro/gorest/src/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// initialize and setup app configuration
	appConfig := config.InitAppConfig()
	// setup server log
	appConfig.SetupAppLog()
	// setup database connection
	appConfig.SetupDatabaseConnection()
	// setup cache client connection
	appConfig.SetupRedisClientConnection()
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	appEngine := gin.Default()
	// register custom middleware
	httpMiddleware := middleware.InitHttpMiddleware()
	// use custom middleware
	appEngine.Use(httpMiddleware.CORS())
	// Initialize data repositories (mysql)
	exampleMySQLRepository := dataSourceMySQL.NewMySQLExampleRepository(
		appConfig.GetDatabaseConnection())
	// Initialize data repository (cache) for cache
	redisCache := dataCache.NewRedisCache(
		appConfig.GetRedisClientConnection(), appConfig.GetCacheTTL())
	// Initialize app use case (service)
	exampleService := useCase.NewExampleService(exampleMySQLRepository)
	// initialize http handler
	httpHandler.NewMainHandler(appEngine, appConfig)
	httpHandler.NewExampleHandler(appEngine, exampleService, redisCache)
	// run the server
	log.Fatal(appEngine.Run(appConfig.GetServerPort()))
}