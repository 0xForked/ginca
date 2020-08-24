package handler

import (
	"fmt"
	httpDelivery "github.com/aasumitro/ginca/src/delivery/http"
	"github.com/aasumitro/ginca/src/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// ExampleHandler represent the http handler for example
type exampleHandler struct {
	exampleService domain.ExampleServiceContract
	exampleCache   domain.RedisCacheContract
}

var handlerName = "examples"

// NewExampleHandler will initialize the example resources endpoint
func NewExampleHandler(
	router *gin.Engine,
	service domain.ExampleServiceContract,
	redis domain.RedisCacheContract,
) {
	handler := &exampleHandler{exampleService: service, exampleCache: redis}
	v1 := router.Group("/v1")
	v1.GET("/examples", handler.fetch)
	v1.GET("/examples/:id", handler.find)
	v1.POST("/examples", handler.create)
	v1.PUT("examples/:id", handler.edit)
	v1.DELETE("/examples/:id", handler.destroy)
}

// Fetch will get all the example data
// FetchExample godoc
// @Summary Retrieves all examples data
// @Produce json
// @Success 200 {array} domain.Example
// @Router /examples [get]
func (handler exampleHandler) fetch(context *gin.Context) {
	var examples = handler.exampleCache.Get(
		context, fmt.Sprintf("%s:all", handlerName))

	if examples == nil {
		examples, err:= handler.exampleService.Fetch()
		if err != nil {
			context.JSON(http.StatusBadRequest, httpDelivery.Respond{
				Code : http.StatusBadRequest,
				Status: err.Error(),
			})
			return
		}

		handler.exampleCache.Set(
			context, fmt.Sprintf("%s:all", handlerName), examples)

		context.JSON(http.StatusOK, httpDelivery.Respond{
			Code : http.StatusOK,
			Status : http.StatusText(http.StatusOK),
			Data : examples,
		})
	} else {
		context.JSON(http.StatusOK, httpDelivery.Respond{
			Code : http.StatusOK,
			Status : http.StatusText(http.StatusOK),
			Data : examples,
		})
	}
}

// Find will get example data by id
// FindExample godoc
// @Summary Retrieves example data based on given ID
// @Produce json
// @Param id path integer true "Example ID"
// @Success 200 {object} domain.Example
// @Router /users/{id} [get]
func (handler exampleHandler) find(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic("error")
	}

	var example = handler.exampleCache.Get(
		context, fmt.Sprintf("%s:%d", handlerName, id))

	if example == nil {
		example, err := handler.exampleService.Find(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, httpDelivery.Respond{
				Code : http.StatusBadRequest,
				Status: err.Error(),
			})
			return
		}

		handler.exampleCache.Set(
			context, fmt.Sprintf("%s:%d", handlerName, id), example)

		context.JSON(http.StatusOK, httpDelivery.Respond{
			Code : http.StatusOK,
			Status : http.StatusText(http.StatusOK),
			Data : example,
		})
	} else {
		context.JSON(http.StatusOK, httpDelivery.Respond{
			Code : http.StatusOK,
			Status : http.StatusText(http.StatusOK),
			Data : example,
		})
	}
}

func (handler exampleHandler) create(context *gin.Context) {
	var example domain.Example

	if err := context.ShouldBind(&example); err != nil {
		context.JSON(http.StatusUnprocessableEntity, httpDelivery.Respond{
			Code : http.StatusUnprocessableEntity,
			Status: err.Error(),
		})
		return
	}

	if err := handler.exampleService.Store(&example); err != nil {
		context.JSON(http.StatusBadRequest, httpDelivery.Respond{
			Code : http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, httpDelivery.Respond{
		Code : http.StatusCreated,
		Status : http.StatusText(http.StatusCreated),
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
		context.JSON(http.StatusUnprocessableEntity, httpDelivery.Respond{
			Code : http.StatusUnprocessableEntity,
			Status: err.Error(),
		})
		return
	}

	example.ID = uint(id)
	example.UpdatedAt = time.Now()

	if err := handler.exampleService.Update(&example); err != nil {
		context.JSON(http.StatusBadRequest, httpDelivery.Respond{
			Code : http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	handler.exampleCache.Delete(
		context, fmt.Sprintf("%s:%d", handlerName, id))

	context.JSON(http.StatusOK, httpDelivery.Respond{
		Code : http.StatusOK,
		Status : http.StatusText(http.StatusOK),
		Data: example,
	})
}

func (handler exampleHandler) destroy(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic("error")
	}

	if err := handler.exampleService.Delete(id); err != nil {
		context.JSON(http.StatusBadRequest, httpDelivery.Respond{
			Code : http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	handler.exampleCache.Delete(
		context, fmt.Sprintf("%s:%d", handlerName, id))

	context.JSON(http.StatusNoContent, httpDelivery.Respond{
		Code : http.StatusNoContent,
		Status : http.StatusText(http.StatusNoContent),
	})
}
