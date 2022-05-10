package handler

import (
	"context"
	"gateway/proto/pb"
	"gateway/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func UserRegister(c *gin.Context) {
	var userReq pb.UserRequest
	err := c.Bind(&userReq)
	if err != nil {
		log.Println(err)
		return
	}
	userService := c.Keys["userService"].(pb.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"data": userResp,
	})
}

func UserLogin(c *gin.Context) {
	var user pb.UserRequest
	err := c.Bind(&user)
	if err != nil {
		log.Println(err)
		return
	}
	userService := c.Keys["userService"].(pb.UserServiceServer)
	userResp, err := userService.UserLogin(context.Background(), &user)
	if err != nil {
		log.Println(err)
		return
	}
	token, err := utils.GenerateToken(userResp.UserDetail.Id)
	c.JSON(200, gin.H{
		"status": 200,
		"msg":    "success",
		"data": gin.H{
			"user":  userResp.UserDetail.UserName,
			"token": token,
		},
	})
}
