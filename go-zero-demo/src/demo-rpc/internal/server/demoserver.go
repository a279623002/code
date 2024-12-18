// Code generated by goctl. DO NOT EDIT.
// Source: demo.proto

package server

import (
	"context"

	"demo-rpc/internal/logic"
	"demo-rpc/internal/svc"
	"demo-rpc/types/demo"
)

type DemoServer struct {
	svcCtx *svc.ServiceContext
	demo.UnimplementedDemoServer
}

func NewDemoServer(svcCtx *svc.ServiceContext) *DemoServer {
	return &DemoServer{
		svcCtx: svcCtx,
	}
}

// rpc方法
func (s *DemoServer) GetID(ctx context.Context, in *demo.DemoRequest) (*demo.DemoResponse, error) {
	l := logic.NewGetIDLogic(ctx, s.svcCtx)
	return l.GetID(in)
}
