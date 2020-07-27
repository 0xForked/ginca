package middleware

import "github.com/gin-gonic/gin"

//Metrics to prometheus
func (middleware HttpMiddleware) Prometheus() gin.HandlerFunc {
	return func(context *gin.Context) {
		// do something
	}
}