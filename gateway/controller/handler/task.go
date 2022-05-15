package handler

import (
	"context"
	"gateway/middleware"
	"gateway/proto/pb"
	"gateway/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CreateTask(c *gin.Context) {
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	var taskReq pb.TaskRequest
	taskReq.Uid = uint64(claims.Id)
	err := c.Bind(&taskReq)
	if err != nil {
		log.Println(err)
		return
	}
	taskService := middleware.MicroService["taskService"].(pb.TaskServiceClient)
	res, err := taskService.CreateTask(context.Background(), &taskReq)
	c.JSON(200, gin.H{
		"data": res,
	})
}

func GetListTask(c *gin.Context) {
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	var taskReq pb.TaskRequest
	taskReq.Uid = uint64(claims.Id)
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		c.JSON(400, err)
	}
	taskReq.Status = int64(status)
	err = c.Bind(&taskReq)
	if err != nil {
		log.Println(err)
		return
	}
	taskService := middleware.MicroService["taskService"].(pb.TaskServiceClient)
	resp, err := taskService.GetListTask(context.Background(), &taskReq)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"data":  resp.TaskList,
		"total": resp.Count,
	})
}

func GetSomeTask(c *gin.Context) {
	var taskReq pb.TaskRequest
	err := c.Bind(&taskReq)
	if err != nil {
		log.Println(err)
		return
	}
	taskService := middleware.MicroService["taskService"].(pb.TaskServiceClient)
	resp, err := taskService.GetSomeTask(context.Background(), &taskReq)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"data": resp,
	})
}

func UpdateTask(c *gin.Context) {
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	var taskReq pb.TaskRequest
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
	}
	taskReq.Id = uint64(id)
	taskReq.Uid = uint64(claims.Id)
	err = c.Bind(&taskReq)
	if err != nil {
		log.Println(err)
		return
	}
	taskService := middleware.MicroService["taskService"].(pb.TaskServiceClient)
	resp, err := taskService.UpdateTask(context.Background(), &taskReq)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"data": resp,
	})
}

func DeleteTask(c *gin.Context) {
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	var taskReq pb.TaskRequest
	taskReq.Uid = uint64(claims.Id)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
	}
	taskReq.Id = uint64(id)
	err = c.Bind(&taskReq)
	if err != nil {
		panic(err)
	}
	taskService := middleware.MicroService["taskService"].(pb.TaskServiceClient)
	resp, err := taskService.DeleteTask(context.Background(), &taskReq)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"data":    resp,
		"message": "删除成功",
	})
}
