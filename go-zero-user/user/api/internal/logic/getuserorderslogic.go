package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-order/order/rpc/types/order"
	"go-zero-user/user/api/internal/svc"
	"go-zero-user/user/api/internal/types"
)

type GetUserOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOrdersLogic {
	return &GetUserOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUserOrders 通过 order-rpc 查询用户订单，user 服务不直接访问订单库
func (l *GetUserOrdersLogic) GetUserOrders(req *types.GetUserOrdersReq) (*types.GetUserOrdersResp, error) {
	// 校验用户存在
	if _, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id); err != nil {
		return nil, err
	}

	resp, err := l.svcCtx.OrderRpc.ListOrder(l.ctx, &order.ListOrderReq{
		UserId:   req.Id,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	list := make([]types.UserOrder, 0, len(resp.List))
	for _, item := range resp.List {
		list = append(list, types.UserOrder{
			Id:          item.Id,
			OrderNo:     item.OrderNo,
			TotalAmount: item.TotalAmount,
			Status:      item.Status,
			Remark:      item.Remark,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	return &types.GetUserOrdersResp{
		List:  list,
		Total: resp.Total,
	}, nil
}
