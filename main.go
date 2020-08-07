package main

import (
	"github.com/aasumitro/ginca/config"
	dataCache "github.com/aasumitro/ginca/src/cache"
	httpHandler "github.com/aasumitro/ginca/src/delivery/http/handler"
	"github.com/aasumitro/ginca/src/delivery/http/middleware"
	dataSourceMySQL "github.com/aasumitro/ginca/src/repository/mysql"
	useCase "github.com/aasumitro/ginca/src/service"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
)

func main() {
	// sets the maximum number of CPUs that can be executing
	runtime.GOMAXPROCS(runtime.NumCPU())
	// initialize and setup app configuration
	appConfig := config.InitAppConfig()
	// load server environment
	appConfig.SetupServerEnvironment()
	// setup server log
	appConfig.SetupServerLog()
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