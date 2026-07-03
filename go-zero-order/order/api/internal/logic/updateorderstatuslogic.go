package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-order/order/api/internal/svc"
	"go-zero-order/order/api/internal/types"
	"go-zero-order/order/rpc/types/order"
)

type UpdateOrderStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(req *types.UpdateOrderStatusReq) (*types.UpdateOrderStatusResp, error) {
	resp, err := l.svcCtx.OrderRpc.UpdateOrderStatus(l.ctx, &order.UpdateOrderStatusReq{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateOrderStatusResp{Success: resp.Success}, nil
}
