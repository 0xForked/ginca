package middleware_test

import (
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

//func TestHttpMiddleware_CORS(t *testing.T) {
//	engine := gin.Default()
//	req := httptest.NewRequest(http.MethodGet, "/",nil)
//	res := httptest.NewRecorder()
//
//	m := middleware.InitHttpMiddleware()
//
//	h := m.CORS(gin.HandlerFunc(func(context *Context)  error {
//		return context.NoContent(http.StatusOK)
//	}))
//}