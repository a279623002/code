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
		user.GET("/registAccount", handler.RegistAccount)
		user.GET("/loginAccount", handler.LoginAccount)
		user.GET("/wantScore", handler.WantScore)
		user.GET("/updateUserProfile", handler.UpdateUserProfile)
	}
	cinema := a.route.Group("/cinema")
	{
		cinema.GET("/locationCinema", handler.LocationCinema)
		cinema.GET("/getCinemaMessageByCid", handler.GetCinemaMessageByCid)
		cinema.GET("/getMovieHallByMHId", handler.GetMovieHallByMHId)
	}
	place := a.route.Group("/place")
	{
		place.GET("/hotCitiesByCinema", handler.HotCitiesByCinema)
	}
	film := a.route.Group("/film")
	{
		film.GET("/hotPlayMovies", handler.HotPlayMovies)
		film.GET("/movieDetail", handler.MovieDetail)
		film.GET("/movieCreditsWithTypes", handler.MovieCreditsWithTypes)
		film.GET("/imageAll", handler.ImageAll)
		film.GET("/locationMovies", handler.LocationMovies)
		film.GET("/movieComingNew", handler.MovieComingNew)
		film.GET("/getFilmsByCidADay", handler.GetFilmsByCidADay)
		film.GET("/search", handler.Search)
	}
	comment := a.route.Group("/comment")
	{
		comment.GET("/hotComment", handler.HotComment)
		comment.GET("/makeComment", handler.MakeComment)
		comment.GET("/upNumComment", handler.UpNumComment)
		comment.GET("/myComments", handler.MyComments)
		comment.GET("/deleteComment", handler.DeleteComment)
	}
	order := a.route.Group("/order")
	{
		order.GET("/WantTicket", handler.WantTicket)
		order.GET("/Ticket", handler.Ticket)
		order.GET("/PayOrder", handler.PayOrder)
		order.GET("/UndoOrder", handler.UndoOrder)
		order.GET("/LookOrders", handler.LookOrders)
		order.GET("/LookAlreadyOrders", handler.LookAlreadyOrders)
		order.GET("/OrderComment", handler.OrderComment)
		order.GET("/GetOrderMessage", handler.GetOrderMessage)
	}
	a.route.Run(port)
}

var DefaultHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "hello",
	})
}
