package main

import (
	"github.com/aasumitro/gorest/config"
	"github.com/aasumitro/gorest/src/http/handler"
	"github.com/aasumitro/gorest/src/http/middleware"
	dataSource "github.com/aasumitro/gorest/src/repository/mysql"
	useCase "github.com/aasumitro/gorest/src/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// initialize and setup app configuration
	appConfig := config.InitAppConfig()
	// setup server access log
	appConfig.SetupAccessLog()
	// setup database connection
	appConfig.SetupDatabaseConnection()
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	appEngine := gin.Default()
	// register custom middleware
	httpMiddleware := middleware.InitHttpMiddleware()
	// use custom middleware
	appEngine.Use(httpMiddleware.CORS())
	// Initialize data repositories
	repository := dataSource.NewMySQLExampleRepository(appConfig.GetDBConnection())
	// Initialize app use case (service)
	service := useCase.NewExampleService(repository)
	// initialize http handler
	handler.NewExampleHandler(appEngine, service)
	// run the server
	log.Fatal(appEngine.Run(appConfig.GetServerPort()))
}