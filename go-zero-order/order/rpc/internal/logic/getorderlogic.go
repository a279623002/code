package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-order/order/rpc/internal/model"
	"go-zero-order/order/rpc/internal/svc"
	"go-zero-order/order/rpc/types/order"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *order.GetOrderReq) (*order.GetOrderResp, error) {
	o, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &order.GetOrderResp{
		Order: toOrderPb(o),
	}, nil
}

func toOrderPb(o *model.Order) *order.Order {
	return &order.Order{
		Id:          o.Id,
		OrderNo:     o.OrderNo,
		UserId:      o.UserId,
		TotalAmount: o.TotalAmount,
		Status:      int32(o.Status),
		Remark:      o.Remark,
		CreatedAt:   o.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   o.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
