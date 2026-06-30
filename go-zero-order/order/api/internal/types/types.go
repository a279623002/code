package types

// Order 订单响应对象
type Order struct {
	Id          int64   `json:"id"`
	OrderNo     string  `json:"order_no"`
	UserId      int64   `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      int32   `json:"status"`
	Remark      string  `json:"remark"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// CreateOrderReq 创建订单请求
type CreateOrderReq struct {
	UserId      int64   `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Remark      string  `json:"remark"`
}

// CreateOrderResp 创建订单响应
type CreateOrderResp struct {
	Id      int64  `json:"id"`
	OrderNo string `json:"order_no"`
}

// GetOrderReq 查询订单请求
type GetOrderReq struct {
	Id int64 `path:"id"`
}

// GetOrderResp 查询订单响应
type GetOrderResp struct {
	Order Order `json:"order"`
}

// ListOrderReq 订单列表请求
type ListOrderReq struct {
	UserId   int64 `form:"user_id,optional"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"page_size,default=20"`
}

// ListOrderResp 订单列表响应
type ListOrderResp struct {
	List  []Order `json:"list"`
	Total int64   `json:"total"`
}

// UpdateOrderStatusReq 更新订单状态请求
type UpdateOrderStatusReq struct {
	Id     int64 `path:"id"`
	Status int32 `json:"status"`
}

// UpdateOrderStatusResp 更新订单状态响应
type UpdateOrderStatusResp struct {
	Success bool `json:"success"`
}
