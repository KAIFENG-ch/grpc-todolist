package handler

import (
	"context"
	"google.golang.org/grpc/stats"
)

type StatsHandler struct {
}

func (h StatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	//TODO implement me
	panic("implement me")
}

func (h StatsHandler) HandleRPC(ctx context.Context, rpcStats stats.RPCStats) {
	//TODO implement me
	panic("implement me")
}

func (h StatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	//TODO implement me
	panic("implement me")
}

func (h StatsHandler) HandleConn(ctx context.Context, connStats stats.ConnStats) {
	//TODO implement me
	panic("implement me")
}
