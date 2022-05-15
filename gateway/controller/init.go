package controller

import (
	"gateway/controller/handler"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	router := gin.Default()
	middleware.RegisterMiddleware(service)
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "SUCCESS")
	})
	router.POST("/user/register", handler.UserRegister)
	router.POST("/user/login", handler.UserLogin)
	v1 := router.Group("/")
	v1.Use(middleware.JWT())
	{
		v1.POST("task", handler.CreateTask)
		v1.GET("tasks/:status", handler.GetListTask)
		v1.GET("tasks", handler.GetSomeTask)
		v1.PUT("task/:id", handler.UpdateTask)
		v1.DELETE("task/:id", handler.DeleteTask)
	}
	return router
}
