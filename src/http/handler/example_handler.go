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
type exampleHandler struct {
	exampleService domain.ExampleService
}

// NewExampleHandler will initialize the example resources endpoint
func NewExampleHandler(router *gin.Engine, service domain.ExampleService) {
	handler := &exampleHandler{exampleService: service}
	v1 := router.Group("/v1")
	v1.GET("/examples", handler.fetch)
	v1.GET("/examples/:id", handler.find)
}

// Fetch will get all the example data
func (handler exampleHandler) fetch(context *gin.Context) {
	examples, _:= handler.exampleService.Fetch()

	context.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"message" : http.StatusText(http.StatusOK),
		"result" : examples,
	})
}

// Find will get example data by id
func (handler exampleHandler) find(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic("error")
	}

	example, _ := handler.exampleService.Find(id)

	if example.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"code" : http.StatusNotFound,
			"message" : http.StatusText(http.StatusNotFound),
			"result": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"message" : http.StatusText(http.StatusOK),
		"result" : example,
	})
}


func (handler exampleHandler) create(context *gin.Context) {

}

func (handler exampleHandler) edit(context *gin.Context) {

}

func (handler exampleHandler) destroy(context *gin.Context) {

}