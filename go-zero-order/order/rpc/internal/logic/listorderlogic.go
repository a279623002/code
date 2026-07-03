package logic

import (
	"context"

	"go-zero-order/order/rpc/internal/svc"
	"go-zero-order/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrderLogic {
	return &ListOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListOrderLogic) ListOrder(in *order.ListOrderReq) (*order.ListOrderResp, error) {
	list, total, err := l.svcCtx.OrderModel.List(l.ctx, in.UserId, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	resp := &order.ListOrderResp{
		Total: total,
		List:  make([]*order.Order, 0, len(list)),
	}
	for _, o := range list {
		resp.List = append(resp.List, toOrderPb(o))
	}
	return resp, nil
}
