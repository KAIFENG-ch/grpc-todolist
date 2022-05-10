package middleware

import "github.com/gin-gonic/gin"

func RegisterMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["userService"] = service[0]
		context.Next()
	}
}
