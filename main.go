package main

import (
	"github.com/aasumitro/gorest/config"
	"github.com/aasumitro/gorest/src/http/handler"
	"github.com/aasumitro/gorest/src/http/middleware"
	dataSourceMySQL "github.com/aasumitro/gorest/src/repository/mysql"
	dataSourceRedis "github.com/aasumitro/gorest/src/repository/redis"
	useCase "github.com/aasumitro/gorest/src/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	// initialize and setup app configuration
	appConfig := config.InitAppConfig()
	// setup server access log when on development mode
	if appConfig.IsDevelopmentMode() {
		appConfig.SetupAccessLog()
	}
	// setup database connection
	appConfig.SetupDatabaseConnection()
	// setup redis client connection
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
	// Initialize data repository (redis) for cache
	redisRepository := dataSourceRedis.NewRedisCache(
		appConfig.GetRedisClientConnection(), time.Minute)
	// Initialize app use case (service)
	exampleService := useCase.NewExampleService(exampleMySQLRepository)
	// initialize http handler
	handler.NewMainHandler(appEngine, redisRepository)
	handler.NewExampleHandler(appEngine, exampleService, redisRepository)
	// run the server
	log.Fatal(appEngine.Run(appConfig.GetServerPort()))
}