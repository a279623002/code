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
	a.route.Run(port)
}

var DefaultHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "hello",
	})
}
