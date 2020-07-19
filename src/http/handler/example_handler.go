package handler

import (
	"github.com/aasumitro/gorest/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ResponseError represent the responses error struct
type Response struct {
	Code 	int64 			`json:"code"`
	Status	string			`json:"status"`
	Data	interface{} 	`json:"data"`
}

// ExampleHandler represent the http handler for example
type ExampleHandler struct {
	ExampleService domain.ExampleService
}

// NewExampleHandler will initialize the example resources endpoint
func NewExampleHandler(router *gin.Engine, service domain.ExampleService) {
	handler := &ExampleHandler{ExampleService: service}
	v1 := router.Group("/v1")
	v1.GET("/examples", handler.Fetch)
	v1.GET("/examples/:id", handler.Find)
}

// Fetch will get all the example data
func (handler ExampleHandler) Fetch(context *gin.Context) {
	examples, _:= handler.ExampleService.Fetch()
	context.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"message" : http.StatusText(http.StatusOK),
		"result" : examples,
	})
}

// Find will get example data by id
func (handler ExampleHandler) Find(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic("error")
	}
	example, _ := handler.ExampleService.Find(id)
	context.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"message" : http.StatusText(http.StatusOK),
		"result" : example,
	})
}
