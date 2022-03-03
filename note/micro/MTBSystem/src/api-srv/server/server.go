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
		order.GET("/wantTicket", handler.WantTicket)
		order.GET("/ticket", handler.Ticket)
		order.GET("/payOrder", handler.PayOrder)
		order.GET("/undoOrder", handler.UndoOrder)
		order.GET("/lookOrders", handler.LookOrders)
		order.GET("/lookAlreadyOrders", handler.LookAlreadyOrders)
		order.GET("/orderComment", handler.OrderComment)
		order.GET("/getOrderMessage", handler.GetOrderMessage)
	}
	cms := a.route.Group("/cms")
	{
		cms.GET("/userLogin", handler.UserLogin)
		cms.GET("/updateMessage", handler.UpdateMessage)
		cms.GET("/allFilms", handler.AllFilms)
		cms.GET("/allUsers", handler.AllUsers)
		cms.GET("/allAdminUsers", handler.AllAdminUsers)
		cms.GET("/allComments", handler.AllComments)
		cms.GET("/allOrders", handler.AllOrders)
		cms.GET("/allAddress", handler.AllAddress)
		cms.GET("/addFilm", handler.AddFilm)
		cms.GET("/updateFilm", handler.UpdateFilm)
		cms.GET("/deleteFilm", handler.DeleteFilm)
		cms.GET("/addAdminUser", handler.AddAdminUser)
		cms.GET("/addAddress", handler.AddAddress)
		cms.GET("/updateAddress", handler.UpdateAddress)
		cms.GET("/deleteAddress", handler.DeleteAddress)
		cms.GET("/deleteAdminUser", handler.DeleteAdminUser)
		cms.GET("/allMovieHall", handler.AllMovieHall)
		cms.GET("/addMovieHall", handler.AddMovieHall)
		cms.GET("/updateMovieHall", handler.UpdateMovieHall)
		cms.GET("/deleteMovieHall", handler.DeleteMovieHall)
		cms.GET("/allCinemaFilms", handler.AllCinemaFilms)
		cms.GET("/addCinemaFilm", handler.AddCinemaFilm)
		cms.GET("/updateCinemaFilm", handler.UpdateCinemaFilm)
		cms.GET("/deleteCinemaFilm", handler.DeleteCinemaFilm)
		cms.GET("/registerCinema", handler.RegisterCinema)
		cms.GET("/allCinemaHall", handler.AllCinemaHall)
	}
	a.route.Run(port)
}

var DefaultHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "hello",
	})
}
