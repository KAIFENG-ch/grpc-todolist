package main

import (
	"context"
	"gateway/controller"
	"gateway/proto/pb"
	"google.golang.org/grpc"
	"time"
)

func main() {
	//etcdResolverBuilder := discover.NewEtcdResolver()
	//resolver.Register(etcdResolverBuilder)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	//taskConn, err := grpc.DialContext(ctx, "127.0.0.1:8001",
	//	grpc.WithStatsHandler(&handler.StatsHandler{}),
	//	grpc.WithInsecure(),
	//	grpc.WithBalancerName(loadBalance.WEIGHT_LOAD_BALANCE),
	//)
	taskConn, err := grpc.DialContext(ctx, "127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//userConn, err := grpc.DialContext(ctx, "127.0.0.1:8002",
	//	grpc.WithStatsHandler(&handler.StatsHandler{}),
	//	grpc.WithInsecure(),
	//	grpc.WithBalancerName(loadBalance.WEIGHT_LOAD_BALANCE),
	//)
	userConn, err := grpc.DialContext(ctx, "127.0.0.1:8002", grpc.WithInsecure())
	taskService := pb.NewTaskServiceClient(taskConn)
	userService := pb.NewUserServiceClient(userConn)
	engine := controller.NewRouter(userService, taskService)
	_ = engine.Run(":8000")
}
