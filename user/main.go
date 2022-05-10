package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user/conf"
	pb "user/proto/proto"
	"user/register"
	"user/service"
)

func main() {
	conf.Init()
	clientReg, err := register.NewEtcdReg()
	if err != nil {
		log.Printf("创建服务错误: %v\n", err)
		panic(err)
	}
	defer func(clientReg *register.Register) {
		err := clientReg.Close()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}(clientReg)
	err = clientReg.RegisterServer("etcdUserService", "127.0.0.1:2379", 5)
	if err != nil {
		log.Printf("注册错误：%v\n", err)
		panic(err)
	}
	lis, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &service.U)
	_ = server.Serve(lis)
}
