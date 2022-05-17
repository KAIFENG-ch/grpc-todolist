package handler

import (
	"fmt"
	"google.golang.org/grpc"
)

func UnknownServiceHandler(srv interface{}, stream grpc.ServerStream) error {
	resp := "服务未找到"
	err := stream.SendMsg(resp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
