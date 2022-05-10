package main

import (
	"context"
	"gateway/controller"
	"gateway/proto/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"time"
)

type Server struct {
	engine    *gin.Engine
	webClient pb.UserServiceClient
}

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	conn, err := grpc.DialContext(ctx, "127.0.0.1:8002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewUserServiceClient(conn)
	s := Server{
		engine:    controller.NewRouter("userService"),
		webClient: client,
	}
	_ = s.engine.Run(":8000")
}
