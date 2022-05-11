package handler

import (
	"context"
	"gateway/proto/pb"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateTask(c *gin.Context) {
	var taskReq pb.TaskRequest
	err := c.Bind(&taskReq)
	if err != nil {
		log.Println(err)
		return
	}
	taskService := c.Keys["taskService"].(pb.TaskServiceClient)
	res, err := taskService.CreateTask(context.Background(), &taskReq)
	c.JSON(200, gin.H{
		"data": res,
	})
}
