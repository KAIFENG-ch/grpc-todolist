package main

import (
	"context"
	"gateway/controller"
	"gateway/proto/pb"
	"google.golang.org/grpc"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	taskConn, err := grpc.DialContext(ctx, "127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	taskService := pb.NewTaskServiceClient(taskConn)
	userConn, err := grpc.DialContext(ctx, "127.0.0.1:8002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userService := pb.NewUserServiceClient(userConn)
	engine := controller.NewRouter(userService, taskService)
	_ = engine.Run(":8000")
}
