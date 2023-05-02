package handler

import (
	"config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"net/http"
	order "order-srv/proto"
	"strconv"
)

var (
	serviceOrder              = config.Namespace + config.ServiceNameOrder
	endpointWantTicket        = "Order.WantTicket"
	endpointTicket            = "Order.Ticket"
	endpointUndoOrder         = "Order.UndoOrder"
	endpointLookOrders        = "Order.LookOrders"
	endpointLookAlreadyOrders = "Order.LookAlreadyOrders"
	endpointOrderComment      = "Order.OrderComment"
	endpointGetOrderMessage   = "Order.GetOrderMessage"
	endpointPayOrder          = "Order.PayOrder"
)

// @Summary 订单详情
// @Tags 影片中心
// @Description 订单详情
// @Accept json
// @Produce json
// @Param userId query int true "用户id"
// @Router /order/getOrderMessage [post]
func GetOrderMessage(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("userId"))
	orderNum := c.Query("orderNum")
	grpcReq := &order.GetOrderMessageReq{
		OrderNum: orderNum,
		UserId:   int64(userId),
	}
	grpcRsp := &order.GetOrderMessageRsp{}

	req := client.NewRequest(serviceOrder, endpointGetOrderMessage, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 订单评分
// @Tags 影片中心
// @Description 订单评分
// @Accept json
// @Produce json
// @Param score query int true "分数"
// @Param userId query int true "用户id"
// @Router /order/orderComment [post]
func OrderComment(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("userId"))
	score, _ := strconv.Atoi(c.Query("score"))
	orderNum := c.Query("orderNum")
	content := c.Query("content")
	grpcReq := &order.OrderCommentReq{
		UserId:         int64(userId),
		Score:          int64(score),
		CommentContent: content,
		OrderNum:       orderNum,
	}
	grpcRsp := &order.OrderCommentRsp{}

	req := client.NewRequest(serviceOrder, endpointOrderComment, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 查看看过的电影
// @Tags 影片中心
// @Description 查看看过的电影
// @Accept json
// @Produce json
// @Param userId query int true "用户id"
// @Router /order/lookAlreadyOrders [post]
func LookAlreadyOrders(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("userId"))
	grpcReq := &order.LookAlreadyOrdersReq{
		UserId: int64(userId),
	}
	grpcRsp := &order.LookAlreadyOrdersRsp{}

	req := client.NewRequest(serviceOrder, endpointLookAlreadyOrders, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 查看所有(订单)电影票
// @Tags 影片中心
// @Description 查看所有(订单)电影票
// @Accept json
// @Produce json
// @Param userId query int true "用户id"
// @Router /order/lookOrders [post]
func LookOrders(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("userId"))
	grpcReq := &order.LookOrdersReq{
		UserId: int64(userId),
	}
	grpcRsp := &order.LookOrdersRsp{}

	req := client.NewRequest(serviceOrder, endpointLookOrders, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

func UndoOrder(c *gin.Context) {
	grpcReq := &order.UndoOrderReq{}
	grpcRsp := &order.UndoOrderRsp{}

	req := client.NewRequest(serviceOrder, endpointUndoOrder, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 付款并更新用户手机
// @Tags 影片中心
// @Description 付款并更新用户手机
// @Accept json
// @Produce json
// @Param userId query int true "用户id"
// @Router /order/payOrder [post]
func PayOrder(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("userId"))
	phone, _ := strconv.Atoi(c.Query("phone"))
	orderNum := c.Query("orderNum")
	grpcReq := &order.PayOrderReq{
		UserId:   int64(userId),
		OrderNum: orderNum,
		Phone:    int64(phone),
	}
	grpcRsp := &order.PayOrderRsp{}

	req := client.NewRequest(serviceOrder, endpointPayOrder, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 选取影厅座位并生成订单
// @Tags 影片中心
// @Description 选取影厅座位并生成订单
// @Accept json
// @Produce json
// @Param mhId query int true "影厅id"
// @Param filmId query int true "影片id"
// @Param userId query int true "用户id"
// @Param x query int true "座位x"
// @Param y query int true "座位y"
// @Param startTime query string true "开始时间"
// @Param endTime query string true "结束时间"
// @Router /order/ticket [post]
func Ticket(c *gin.Context) {
	filmId, _ := strconv.Atoi(c.Query("filmId"))
	userId, _ := strconv.Atoi(c.Query("userId"))
	mhId, _ := strconv.Atoi(c.Query("mhId"))
	x, _ := strconv.Atoi(c.Query("x"))
	y, _ := strconv.Atoi(c.Query("y"))
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	grpcReq := &order.TicketReq{
		UserId:    int64(userId),
		FilmId:    int64(filmId),
		MhId:      int64(mhId),
		X:         int64(x),
		Y:         int64(y),
		StartTime: startTime,
		EndTime:   endTime,
	}
	grpcRsp := &order.TicketRsp{}

	req := client.NewRequest(serviceOrder, endpointTicket, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 记录想看的电影
// @Tags 影片中心
// @Description 记录想看的电影
// @Accept json
// @Produce json
// @Param filmId query int true "影片id"
// @Param userId query int true "用户id"
// @Router /order/wantTicket [post]
func WantTicket(c *gin.Context) {
	filmId, _ := strconv.Atoi(c.Query("filmId"))
	userId, _ := strconv.Atoi(c.Query("userId"))
	grpcReq := &order.WantTicketReq{
		FilmId: int64(filmId),
		UserId: int64(userId),
	}
	grpcRsp := &order.WantTicketRsp{}

	req := client.NewRequest(serviceOrder, endpointWantTicket, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}
