package orderclient

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-service/order/rpc/types/order"
)

// Order 暴露给 api 层使用的 rpc 客户端
type Order interface {
	CreateOrder(ctx context.Context, in *order.CreateOrderReq, opts ...interface{}) (*order.CreateOrderResp, error)
	GetOrder(ctx context.Context, in *order.GetOrderReq, opts ...interface{}) (*order.GetOrderResp, error)
	ListOrder(ctx context.Context, in *order.ListOrderReq, opts ...interface{}) (*order.ListOrderResp, error)
	UpdateOrderStatus(ctx context.Context, in *order.UpdateOrderStatusReq, opts ...interface{}) (*order.UpdateOrderStatusResp, error)
}

// Client 订单 rpc 客户端包装
type Client struct {
	cli order.OrderRpcClient
}

// NewClient 创建客户端
func NewClient(c zrpc.RpcClientConf) (*Client, error) {
	cli, err := zrpc.NewClient(c)
	if err != nil {
		return nil, err
	}
	return &Client{cli: order.NewOrderRpcClient(cli.Conn())}, nil
}

func (c *Client) CreateOrder(ctx context.Context, in *order.CreateOrderReq, opts ...interface{}) (*order.CreateOrderResp, error) {
	return c.cli.CreateOrder(ctx, in)
}

func (c *Client) GetOrder(ctx context.Context, in *order.GetOrderReq, opts ...interface{}) (*order.GetOrderResp, error) {
	return c.cli.GetOrder(ctx, in)
}

func (c *Client) ListOrder(ctx context.Context, in *order.ListOrderReq, opts ...interface{}) (*order.ListOrderResp, error) {
	return c.cli.ListOrder(ctx, in)
}

func (c *Client) UpdateOrderStatus(ctx context.Context, in *order.UpdateOrderStatusReq, opts ...interface{}) (*order.UpdateOrderStatusResp, error) {
	return c.cli.UpdateOrderStatus(ctx, in)
}
