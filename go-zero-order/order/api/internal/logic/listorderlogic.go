package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-order/order/api/internal/svc"
	"go-zero-order/order/api/internal/types"
	"go-zero-order/order/rpc/types/order"
)

type ListOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrderLogic {
	return &ListOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListOrderLogic) ListOrder(req *types.ListOrderReq) (*types.ListOrderResp, error) {
	resp, err := l.svcCtx.OrderRpc.ListOrder(l.ctx, &order.ListOrderReq{
		UserId:   req.UserId,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	list := make([]types.Order, 0, len(resp.List))
	for _, item := range resp.List {
		list = append(list, types.Order{
			Id:          item.Id,
			OrderNo:     item.OrderNo,
			UserId:      item.UserId,
			TotalAmount: item.TotalAmount,
			Status:      item.Status,
			Remark:      item.Remark,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	return &types.ListOrderResp{
		List:  list,
		Total: resp.Total,
	}, nil
}
