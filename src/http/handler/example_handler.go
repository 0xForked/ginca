package handler

import (
	"github.com/aasumitro/gorest/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// ResponseError represent the responses error struct
type response struct {
	Code 	int64 			`json:"code"`
	Message	string			`json:"message"`
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
	v1.POST("/examples", handler.create)
	v1.PUT("examples/:id", handler.edit)
	v1.DELETE("/examples/:id", handler.destroy)
}

// Fetch will get all the example data
func (handler exampleHandler) fetch(context *gin.Context) {
	examples, _:= handler.exampleService.Fetch()

	context.JSON(http.StatusOK, response{
		Code : http.StatusOK,
		Message : http.StatusText(http.StatusOK),
		Data : examples,
	})
}

// Find will get example data by id
func (handler exampleHandler) find(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic("error")
	}

	example, err := handler.exampleService.Find(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, response{
			Code : http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response{
		Code : http.StatusOK,
		Message : http.StatusText(http.StatusOK),
		Data : example,
	})
}

func (handler exampleHandler) create(context *gin.Context) {
	var example domain.Example

	if err := context.ShouldBind(&example); err != nil {
		context.JSON(http.StatusUnprocessableEntity, response{
			Code : http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	if err := handler.exampleService.Store(&example); err != nil {
		context.JSON(http.StatusBadRequest, response{
			Code : http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, response{
		Code : http.StatusCreated,
		Message : http.StatusText(http.StatusCreated),
		Data : example,
	})
}

func (handler exampleHandler) edit(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic("error")
	}

	var example domain.Example

	if err := context.ShouldBind(&example); err != nil {
		context.JSON(http.StatusUnprocessableEntity, response{
			Code : http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	example.ID = uint(id)
	example.UpdatedAt = time.Now()

	if err := handler.exampleService.Update(&example); err != nil {
		context.JSON(http.StatusBadRequest, response{
			Code : http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response{
		Code : http.StatusOK,
		Message : http.StatusText(http.StatusOK),
		Data: example,
	})
}

func (handler exampleHandler) destroy(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic("error")
	}

	if err := handler.exampleService.Delete(id); err != nil {
		context.JSON(http.StatusBadRequest, response{
			Code : http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusNoContent, response{
		Code : http.StatusNoContent,
		Message : http.StatusText(http.StatusNoContent),
	})
}
