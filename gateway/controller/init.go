package controller

import (
	"gateway/controller/handler"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.JWT())
	v1 := router.Group("api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "SUCCESS")
		})
		v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", handler.UserLogin)
	}
	return router
}
