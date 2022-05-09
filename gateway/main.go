package main

import (
	"context"
	"fmt"
	"gateway/discover"
	"gateway/proto/pb"
	"gateway/service"
	"google.golang.org/grpc"
	healthPb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/resolver"
	"time"
)

func main() {
	etcdResolveBuilder := discover.NewEtcdResolver()
	resolver.Register(etcdResolveBuilder)
	etcdReg, err := service.NewEtcdReg()
	err = etcdReg.RegisterServer("userServiceClient", "127.0.0.1:2379", 5)
	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	healthClient := healthPb.NewHealthClient(conn)
	ir := &healthPb.HealthCheckRequest{
		Service: "grpc.health.v1.health",
	}
	deadline, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	healthCheckResp, err := healthClient.Check(deadline, ir)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(healthCheckResp)
	cancelFunc()
	client := pb.NewUserServiceClient(conn)
	loginReq := new(pb.UserRequest)
	loginResp, err := client.UserLogin(context.Background(), loginReq)
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	fmt.Println("连接成功", loginResp)
	registerReq := new(pb.UserRequest)
	registerResp, err := client.UserRegister(context.Background(), registerReq)
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	fmt.Println("连接成功", registerResp)
}
