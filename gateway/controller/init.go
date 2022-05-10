package controller

import (
	"gateway/controller/handler"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.RegisterMiddleware(service))
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "SUCCESS")
	})
	router.POST("/user/register", handler.UserRegister)
	router.POST("/user/login", handler.UserLogin)
	return router
}
