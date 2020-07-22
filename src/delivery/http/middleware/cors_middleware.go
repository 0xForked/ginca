package middleware

import "github.com/gin-gonic/gin"

// CORS will handle the CORS middleware
func (middleware HttpMiddleware) CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Next()
	}
}