package server

import (
	"api-srv/server/handler"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

type Api struct {
	client *client.Client
	route  *gin.Engine
}

func NewApi() *Api {
	service := micro.NewService()
	service.Init()
	c := service.Client()
	r := gin.Default()
	return &Api{
		client: &c,
		route:  r,
	}
}

func (a *Api) Run(port string) {
	a.route.GET("/", DefaultHandler)

	user := a.route.Group("/user")
	{
		user.POST("/registAccount", handler.RegistAccount) // 注册
		user.POST("/loginAccount", handler.LoginAccount) // 登录
		user.POST("/wantScore", handler.WantScore) // 评分（需要下单后才能评分）
		user.POST("/updateUserProfile", handler.UpdateUserProfile) // 更新
	}
	cinema := a.route.Group("/cinema")
	{
		cinema.POST("/locationCinema", handler.LocationCinema) // 根据location_id 获取影院
		cinema.POST("/getCinemaMessageByCid", handler.GetCinemaMessageByCid) // 正在销售的影片信息和影院信息
		cinema.POST("/getMovieHallByMHId", handler.GetMovieHallByMHId) // 获取影厅座位表
	}
	place := a.route.Group("/place")
	{
		place.GET("/hotCitiesByCinema", handler.HotCitiesByCinema) //无操作
	}
	film := a.route.Group("/film")
	{
		film.GET("/hotPlayMovies", handler.HotPlayMovies) // 正在售票列表 无操作
		film.POST("/movieDetail", handler.MovieDetail) // 详情
		film.POST("/movieCreditsWithTypes", handler.MovieCreditsWithTypes) // 导演演员
		film.POST("/imageAll", handler.ImageAll) // 剧照
		film.GET("/locationMovies", handler.LocationMovies) // 上映的影片
		film.POST("/movieComingNew", handler.MovieComingNew) // 即将上映的影片
		film.POST("/getFilmsByCidADay", handler.GetFilmsByCidADay) // 获取正在销售的影片信息
		film.GET("/search", handler.Search) // 搜索--暂时写死
	}
	comment := a.route.Group("/comment")
	{
		comment.POST("/hotComment", handler.HotComment) // 获取评论
		comment.POST("/makeComment", handler.MakeComment) // 评论
		comment.POST("/upNumComment", handler.UpNumComment) // 评论点赞
		comment.POST("/myComments", handler.MyComments) // 我的评论
		comment.POST("/deleteComment", handler.DeleteComment) // 删除评论
	}
	order := a.route.Group("/order")
	{
		order.POST("/wantTicket", handler.WantTicket) // 记录想看的电影
		order.POST("/ticket", handler.Ticket) // 选取影厅座位并生成订单
		order.POST("/payOrder", handler.PayOrder) // 付款并更新用户手机
		order.GET("/undoOrder", handler.UndoOrder) //无操作
		order.POST("/lookOrders", handler.LookOrders) // 查看所有(订单)电影票
		order.POST("/lookAlreadyOrders", handler.LookAlreadyOrders) // 查看看过的电影
		order.POST("/orderComment", handler.OrderComment) // 订单评分
		order.POST("/getOrderMessage", handler.GetOrderMessage) // 订单详情
	}
	// 管理admin，user，order
	cms := a.route.Group("/cms")
	{
		cms.POST("/userLogin", handler.UserLogin) // 管理员登录
		cms.POST("/updateMessage", handler.UpdateMessage) //无操作
		cms.POST("/allFilms", handler.AllFilms) // 获取所有影片
		cms.POST("/allUsers", handler.AllUsers) // 获取所有用户
		cms.POST("/allAdminUsers", handler.AllAdminUsers) // 获取所有管理员
		cms.POST("/allComments", handler.AllComments) // 获取电影评论(超级管理员1获取所有)
		cms.POST("/allOrders", handler.AllOrders) // 所有订单
		cms.POST("/allAddress", handler.AllAddress) // 获取所有地址
		cms.POST("/addFilm", handler.AddFilm) // 添加影片
		cms.POST("/updateFilm", handler.UpdateFilm) // 更新影片
		cms.POST("/deleteFilm", handler.DeleteFilm) // 删除影片
		cms.POST("/addAdminUser", handler.AddAdminUser) // 添加管理员
		cms.POST("/addAddress", handler.AddAddress) // 添加地址
		cms.POST("/updateAddress", handler.UpdateAddress) // 更新地址
		cms.POST("/deleteAddress", handler.DeleteAddress) // 删除地址
		cms.POST("/deleteAdminUser", handler.DeleteAdminUser) // 删除管理员
		cms.POST("/allMovieHall", handler.AllMovieHall) // 放映厅信息
		cms.POST("/addMovieHall", handler.AddMovieHall) // 添加放映厅
		cms.POST("/updateMovieHall", handler.UpdateMovieHall) // 更新放映厅
		cms.POST("/deleteMovieHall", handler.DeleteMovieHall) // 删除放映厅
		cms.POST("/allCinemaFilms", handler.AllCinemaFilms) // 获取电影院所有影片
		cms.POST("/addCinemaFilm", handler.AddCinemaFilm) // 电影院添加影片
		cms.POST("/updateCinemaFilm", handler.UpdateCinemaFilm) // 更新电影院影片
		cms.POST("/deleteCinemaFilm", handler.DeleteCinemaFilm) // 删除电影院影片
		cms.POST("/registerCinema", handler.RegisterCinema) // 添加电影院
		cms.POST("/allCinemaHall", handler.AllCinemaHall) // 获取电影院所有放映厅
	}
	a.route.Run(port)
}

var DefaultHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "hello",
	})
}
