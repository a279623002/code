package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-service/order/rpc/internal/model"
	"go-zero-service/order/rpc/internal/svc"
	"go-zero-service/order/rpc/types/order"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *order.UpdateOrderStatusReq) (*order.UpdateOrderStatusResp, error) {
	if in.Id <= 0 {
		return nil, fmt.Errorf("invalid id")
	}
	if in.Status < 1 || in.Status > 5 {
		return nil, fmt.Errorf("invalid status")
	}

	if err := l.svcCtx.OrderModel.UpdateStatus(l.ctx, in.Id, model.OrderStatus(in.Status)); err != nil {
		return &order.UpdateOrderStatusResp{Success: false}, err
	}

	return &order.UpdateOrderStatusResp{Success: true}, nil
}
