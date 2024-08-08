package logic

import (
	"context"

	"demo-rpc/internal/svc"
	"demo-rpc/types/demo"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIDLogic {
	return &GetIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc方法
func (l *GetIDLogic) GetID(in *demo.DemoRequest) (*demo.DemoResponse, error) {
	out := &demo.DemoResponse{
		Id: in.Id,
	}

	return out, nil
}
