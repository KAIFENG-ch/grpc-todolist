package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"task/conf"
	"task/proto/pb"
	"task/register"
	"task/service"
)

func main() {
	conf.Init()
	clientReg, err := register.NewEtcdReg()
	if err != nil {
		log.Printf("创建etcd错误:%v\n", err)
		panic(err)
	}
	defer func(clientReg *register.Register) {
		err := clientReg.Close()
		if err != nil {
			log.Println(err)
		}
	}(clientReg)
	err = clientReg.RegisterServer("etcdTaskService", "127.0.0.1:2379", 5)
	if err != nil {
		log.Printf("注册服务失败：%v\n", err)
		panic(err)
	}
	lis, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterTaskServiceServer(server, &service.T)
	_ = server.Serve(lis)
}
