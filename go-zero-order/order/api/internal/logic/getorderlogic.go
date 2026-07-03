package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-order/order/api/internal/svc"
	"go-zero-order/order/api/internal/types"
	"go-zero-order/order/rpc/types/order"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.GetOrderReq) (*types.GetOrderResp, error) {
	resp, err := l.svcCtx.OrderRpc.GetOrder(l.ctx, &order.GetOrderReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetOrderResp{
		Order: types.Order{
			Id:          resp.Order.Id,
			OrderNo:     resp.Order.OrderNo,
			UserId:      resp.Order.UserId,
			TotalAmount: resp.Order.TotalAmount,
			Status:      resp.Order.Status,
			Remark:      resp.Order.Remark,
			CreatedAt:   resp.Order.CreatedAt,
			UpdatedAt:   resp.Order.UpdatedAt,
		},
	}, nil
}
