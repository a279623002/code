package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-service/order/rpc/internal/model"
	"go-zero-service/order/rpc/internal/svc"
	"go-zero-service/order/rpc/types/order"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	if in.UserId <= 0 {
		return nil, fmt.Errorf("invalid user_id")
	}
	if in.TotalAmount <= 0 {
		return nil, fmt.Errorf("invalid total_amount")
	}

	o := &model.Order{
		OrderNo:     generateOrderNo(),
		UserId:      in.UserId,
		TotalAmount: in.TotalAmount,
		Status:      model.OrderStatusPending,
		Remark:      in.Remark,
	}

	if err := l.svcCtx.OrderModel.Create(l.ctx, o); err != nil {
		return nil, err
	}

	return &order.CreateOrderResp{
		Id:      o.Id,
		OrderNo: o.OrderNo,
	}, nil
}

// generateOrderNo 生成订单号：时间前缀 + 毫秒时间戳 + 随机数
func generateOrderNo() string {
	now := time.Now()
	return fmt.Sprintf("ORD%s%03d", now.Format("20060102150405"), now.UnixMilli()%1000)
}
