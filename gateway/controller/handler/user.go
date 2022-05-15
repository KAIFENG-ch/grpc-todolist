package handler

import (
	"context"
	"gateway/middleware"
	"gateway/proto/pb"
	"gateway/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func UserRegister(c *gin.Context) {
	var userReq pb.UserRequest
	err := c.ShouldBind(&userReq)
	if err != nil {
		log.Println(err)
		return
	}
	userService := middleware.MicroService["userService"].(pb.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	if err != nil {
		c.JSON(200, gin.H{
			"data": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": userResp,
	})
}

func UserLogin(c *gin.Context) {
	var user pb.UserRequest
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}
	userService := middleware.MicroService["userService"].(pb.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &user)
	if err != nil {
		panic(err)
	}
	token, err := utils.GenerateToken(userResp.UserDetail.Id)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"data":  userResp,
		"token": token,
	})
}
