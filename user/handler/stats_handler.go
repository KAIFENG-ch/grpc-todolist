package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/stats"
	"log"
)

type StatsHandler struct {
}

func (h StatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	fmt.Println("tagRPC...@" + info.FullMethodName)
	return ctx
}

func (h StatsHandler) HandleRPC(ctx context.Context, rpcStats stats.RPCStats) {
	switch rpcStats.(type) {
	case *stats.Begin:
		fmt.Println("handlerRPC begin...")
	case *stats.End:
		fmt.Println("handlerRPC End...")
	case *stats.InHeader:
		fmt.Println("handlerRPC InHeader...")
	case *stats.InPayload:
		fmt.Println("handlerRPC InPayload...")
	case *stats.InTrailer:
		fmt.Println("handlerRPC InTrailer...")
	case *stats.OutHeader:
		fmt.Println("handlerRPC OutHeader...")
	case *stats.OutPayload:
		fmt.Println("handlerRPC OutPayload...")
	default:
		fmt.Println("handleRPC...")
	}
}

func (h StatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	fmt.Println("tagConn...")
	return ctx
}

func (h StatsHandler) HandleConn(ctx context.Context, connStats stats.ConnStats) {
	switch connStats.(type) {
	case *stats.ConnBegin:
		log.Printf("begin conn")
	case *stats.ConnEnd:
		log.Printf("end conn")
	default:
		fmt.Println("handleConn...")
	}
}
