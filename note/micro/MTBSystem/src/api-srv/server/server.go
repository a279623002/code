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
		user.POST("/registAccount", handler.RegistAccount)
		user.POST("/loginAccount", handler.LoginAccount)
		user.POST("/wantScore", handler.WantScore)
		user.POST("/updateUserProfile", handler.UpdateUserProfile)
	}
	cinema := a.route.Group("/cinema")
	{
		cinema.POST("/locationCinema", handler.LocationCinema)
		cinema.POST("/getCinemaMessageByCid", handler.GetCinemaMessageByCid)
		cinema.POST("/getMovieHallByMHId", handler.GetMovieHallByMHId)
	}
	place := a.route.Group("/place")
	{
		place.GET("/hotCitiesByCinema", handler.HotCitiesByCinema)
	}
	film := a.route.Group("/film")
	{
		film.GET("/hotPlayMovies", handler.HotPlayMovies)
		film.POST("/movieDetail", handler.MovieDetail)
		film.POST("/movieCreditsWithTypes", handler.MovieCreditsWithTypes)
		film.POST("/imageAll", handler.ImageAll)
		film.GET("/locationMovies", handler.LocationMovies)
		film.POST("/movieComingNew", handler.MovieComingNew)
		film.POST("/getFilmsByCidADay", handler.GetFilmsByCidADay)
		film.GET("/search", handler.Search)
	}
	comment := a.route.Group("/comment")
	{
		comment.POST("/hotComment", handler.HotComment)
		comment.POST("/makeComment", handler.MakeComment)
		comment.POST("/upNumComment", handler.UpNumComment)
		comment.POST("/myComments", handler.MyComments)
		comment.POST("/deleteComment", handler.DeleteComment)
	}
	order := a.route.Group("/order")
	{
		order.POST("/wantTicket", handler.WantTicket)
		order.POST("/ticket", handler.Ticket)
		order.POST("/payOrder", handler.PayOrder)
		order.GET("/undoOrder", handler.UndoOrder)
		order.POST("/lookOrders", handler.LookOrders)
		order.POST("/lookAlreadyOrders", handler.LookAlreadyOrders)
		order.POST("/orderComment", handler.OrderComment)
		order.POST("/getOrderMessage", handler.GetOrderMessage)
	}
	cms := a.route.Group("/cms")
	{
		cms.POST("/userLogin", handler.UserLogin)
		cms.POST("/updateMessage", handler.UpdateMessage)
		cms.POST("/allFilms", handler.AllFilms)
		cms.POST("/allUsers", handler.AllUsers)
		cms.POST("/allAdminUsers", handler.AllAdminUsers)
		cms.POST("/allComments", handler.AllComments)
		cms.POST("/allOrders", handler.AllOrders)
		cms.POST("/allAddress", handler.AllAddress)
		cms.POST("/addFilm", handler.AddFilm)
		cms.POST("/updateFilm", handler.UpdateFilm)
		cms.POST("/deleteFilm", handler.DeleteFilm)
		cms.POST("/addAdminUser", handler.AddAdminUser)
		cms.POST("/addAddress", handler.AddAddress)
		cms.POST("/updateAddress", handler.UpdateAddress)
		cms.POST("/deleteAddress", handler.DeleteAddress)
		cms.POST("/deleteAdminUser", handler.DeleteAdminUser)
		cms.POST("/allMovieHall", handler.AllMovieHall)
		cms.POST("/addMovieHall", handler.AddMovieHall)
		cms.POST("/updateMovieHall", handler.UpdateMovieHall)
		cms.POST("/deleteMovieHall", handler.DeleteMovieHall)
		cms.POST("/allCinemaFilms", handler.AllCinemaFilms)
		cms.POST("/addCinemaFilm", handler.AddCinemaFilm)
		cms.POST("/updateCinemaFilm", handler.UpdateCinemaFilm)
		cms.POST("/deleteCinemaFilm", handler.DeleteCinemaFilm)
		cms.POST("/registerCinema", handler.RegisterCinema)
		cms.POST("/allCinemaHall", handler.AllCinemaHall)
	}
	a.route.Run(port)
}

var DefaultHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "hello",
	})
}
