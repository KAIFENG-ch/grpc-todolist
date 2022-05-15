package main

import (
	"context"
	"fmt"
	"gateway/controller"
	"gateway/discover"
	"gateway/handler"
	"gateway/proto/pb"
	"google.golang.org/grpc"
	healthPb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/resolver"
	"time"
)

func main() {
	etcdResolverBuilder := discover.NewEtcdResolver()
	resolver.Register(etcdResolverBuilder)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	taskConn, err := grpc.DialContext(ctx, "127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	taskService := pb.NewTaskServiceClient(taskConn)
	userConn, err := grpc.DialContext(ctx, "127.0.0.1:8002",
		grpc.WithStatsHandler(&handler.StatsHandler{}),
		grpc.WithInsecure(),
		grpc.WithBalancerName("weight_load_balance"),
	)
	if err != nil {
		panic(err)
	}
	healthClient := healthPb.NewHealthClient(userConn)
	ir := &healthPb.HealthCheckRequest{
		Service: "grpc.health.v1.Health",
	}
	deadline, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	healthCheckResp, err := healthClient.Check(deadline, ir)
	if err != nil {
		panic(err)
	}
	fmt.Println(healthCheckResp)
	cancel()
	userService := pb.NewUserServiceClient(userConn)
	engine := controller.NewRouter(userService, taskService)
	_ = engine.Run(":8000")
}
