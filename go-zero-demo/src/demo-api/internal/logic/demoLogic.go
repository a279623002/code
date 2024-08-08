package logic

import (
	"context"
	"demo-rpc/democlient"

	"demo-api/internal/svc"
	"demo-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取id
func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.DemoReq) (resp *types.DemoResp, err error) {
	demoRes, err := l.svcCtx.DemoClient.GetID(l.ctx, &democlient.DemoRequest{Id: req.ID})
	if err != nil {
		return nil, err
	}
	if demoRes == nil {
		return nil, err
	}
	resp = &types.DemoResp{
		ID: demoRes.Id,
	}
	return
}
