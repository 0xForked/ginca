package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ExampleHandler represent the http handler for example
type exampleWSHandler struct {
	// some thing good stuff here
}

//var handlerName = "examples"

// NewExampleHandler will initialize the example resources endpoint
func NewExampleWSHandler(
	router *gin.Engine,
) {
	handler := &exampleWSHandler{}
	router.GET("/ws-pub-example", handler.push)
}

var wsUpgraded = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Fetch will get all the example data
func (handler exampleWSHandler) push(context *gin.Context) {
	conn, err := wsUpgraded.Upgrade(context.Writer,
		context.Request, nil)

	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		if err := conn.WriteMessage(t, msg); err != nil {
			fmt.Printf("Failed to write message: %+v\n", err)
		}
	}
}

