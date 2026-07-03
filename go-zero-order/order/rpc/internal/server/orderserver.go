package server

import (
	"context"

	"go-zero-order/order/rpc/internal/logic"
	"go-zero-order/order/rpc/internal/svc"
	"go-zero-order/order/rpc/types/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderRpcServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) CreateOrder(ctx context.Context, in *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

func (s *OrderServer) GetOrder(ctx context.Context, in *order.GetOrderReq) (*order.GetOrderResp, error) {
	l := logic.NewGetOrderLogic(ctx, s.svcCtx)
	return l.GetOrder(in)
}

func (s *OrderServer) ListOrder(ctx context.Context, in *order.ListOrderReq) (*order.ListOrderResp, error) {
	l := logic.NewListOrderLogic(ctx, s.svcCtx)
	return l.ListOrder(in)
}

func (s *OrderServer) UpdateOrderStatus(ctx context.Context, in *order.UpdateOrderStatusReq) (*order.UpdateOrderStatusResp, error) {
	l := logic.NewUpdateOrderStatusLogic(ctx, s.svcCtx)
	return l.UpdateOrderStatus(in)
}
