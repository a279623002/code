package types

// User 用户对象
type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateUserReq 创建用户请求
type CreateUserReq struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// CreateUserResp 创建用户响应
type CreateUserResp struct {
	Id int64 `json:"id"`
}

// GetUserReq 查询用户请求
type GetUserReq struct {
	Id int64 `path:"id"`
}

// GetUserResp 查询用户响应
type GetUserResp struct {
	User User `json:"user"`
}

// UserOrder 用户视角订单对象
type UserOrder struct {
	Id          int64   `json:"id"`
	OrderNo     string  `json:"order_no"`
	TotalAmount float64 `json:"total_amount"`
	Status      int32   `json:"status"`
	Remark      string  `json:"remark"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// GetUserOrdersReq 用户订单列表请求
type GetUserOrdersReq struct {
	Id       int64 `path:"id"`
	Page     int32 `form:"page,default=1"`
	PageSize int32 `form:"page_size,default=20"`
}

// GetUserOrdersResp 用户订单列表响应
type GetUserOrdersResp struct {
	List  []UserOrder `json:"list"`
	Total int64       `json:"total"`
}

// PlaceOrderReq 下单请求
type PlaceOrderReq struct {
	Id          int64   `path:"id"`
	TotalAmount float64 `json:"total_amount"`
	Remark      string  `json:"remark"`
}

// PlaceOrderResp 下单响应
type PlaceOrderResp struct {
	Id      int64  `json:"id"`
	OrderNo string `json:"order_no"`
}
