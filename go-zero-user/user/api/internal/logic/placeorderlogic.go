package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-order/order/rpc/types/order"
	"go-zero-user/user/api/internal/svc"
	"go-zero-user/user/api/internal/types"
)

type PlaceOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceOrderLogic {
	return &PlaceOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PlaceOrder 通过 order-rpc 为用户下单，user 服务不直接操作订单表
func (l *PlaceOrderLogic) PlaceOrder(req *types.PlaceOrderReq) (*types.PlaceOrderResp, error) {
	// 校验用户存在
	if _, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id); err != nil {
		return nil, err
	}

	resp, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &order.CreateOrderReq{
		UserId:      req.Id,
		TotalAmount: req.TotalAmount,
		Remark:      req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &types.PlaceOrderResp{
		Id:      resp.Id,
		OrderNo: resp.OrderNo,
	}, nil
}
