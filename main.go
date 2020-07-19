package main

import (
	"github.com/aasumitro/gorest/config"
	"github.com/aasumitro/gorest/src/http/handler"
	"github.com/aasumitro/gorest/src/http/middleware"
	"github.com/aasumitro/gorest/src/repository/mysql"
	"github.com/aasumitro/gorest/src/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// initialize and setup app configuration
	appConfig := config.SetupAppConfig()
	// setup database connection
	appConfig.SetupDatabaseConnection()
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	app := gin.Default()
	// register custom middleware
	httpMiddleware := middleware.InitHttpMiddleware()
	// use custom middleware
	app.Use(httpMiddleware.CORS())
	// Initialize data repositories
	repository := mysql.NewMySQLExampleRepository(appConfig.GetDBConnection())
	// Initialize app use case (service)
	useCase := service.NewExampleService(repository)
	// initialize http handler
	handler.NewExampleHandler(app, useCase)
	// run the server
	log.Fatal(app.Run(appConfig.GetServerPort()))
}