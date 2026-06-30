package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-service/order/api/internal/svc"
	"go-zero-service/order/api/internal/types"
	"go-zero-service/order/rpc/types/order"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (*types.CreateOrderResp, error) {
	resp, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &order.CreateOrderReq{
		UserId:      req.UserId,
		TotalAmount: req.TotalAmount,
		Remark:      req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateOrderResp{
		Id:      resp.Id,
		OrderNo: resp.OrderNo,
	}, nil
}
