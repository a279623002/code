package handler

import (
	"config"
	"context"
	film "film-srv/proto"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"net/http"
	"strconv"
)

var (
	serviceFilm                       = config.Namespace + config.ServiceNameFilm
	endpointSearch                = "Film.Search"
	endpointGetFilmsByCidADay     = "Film.GetFilmsByCidADay"
	endpointMovieComingNew        = "Film.MovieComingNew"
	endpointLocationMovies        = "Film.LocationMovies"
	endpointImageAll              = "Film.ImageAll"
	endpointMovieCreditsWithTypes = "Film.MovieCreditsWithTypes"
	endpointMovieDetail           = "Film.MovieDetail"
	endpointHotPlayMovies         = "Film.HotPlayMovies"
)

func HotPlayMovies(c *gin.Context) {
	grpcReq := &film.HotPlayMoviesReq{}
	grpcRsp := &film.HotPlayMoviesRep{}

	req := client.NewRequest(serviceFilm, endpointHotPlayMovies, grpcReq)

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

func MovieDetail(c *gin.Context) {
	movieId, _ := strconv.Atoi(c.Query("movieId"))
	grpcReq := &film.MovieDetailReq{
		MovieId: int64(movieId),
	}
	grpcRsp := &film.MovieDetailRep{}

	req := client.NewRequest(serviceFilm, endpointMovieDetail, grpcReq)

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

func MovieCreditsWithTypes(c *gin.Context) {
	movieId, _ := strconv.Atoi(c.Query("movieId"))
	grpcReq := &film.MovieCreditsWithTypesReq{
		MovieId: int64(movieId),
	}
	grpcRsp := &film.MovieCreditsWithTypesRep{}

	req := client.NewRequest(serviceFilm, endpointMovieCreditsWithTypes, grpcReq)

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

func ImageAll(c *gin.Context) {
	movieId, _ := strconv.Atoi(c.Query("movieId"))
	grpcReq := &film.ImageAllReq{
		MovieId: int64(movieId),
	}
	grpcRsp := &film.ImageAllRep{}

	req := client.NewRequest(serviceFilm, endpointImageAll, grpcReq)

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

func LocationMovies(c *gin.Context) {
	grpcReq := &film.LocationMoviesReq{}
	grpcRsp := &film.LocationMoviesRep{}

	req := client.NewRequest(serviceFilm, endpointLocationMovies, grpcReq)

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

func MovieComingNew(c *gin.Context) {

	grpcReq := &film.GetFilmsByCidADayReq{}
	grpcRsp := &film.GetFilmsByCidADayRsp{}

	req := client.NewRequest(serviceFilm, endpointMovieComingNew, grpcReq)

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

func GetFilmsByCidADay(c *gin.Context) {
	cinemaId, _ := strconv.Atoi(c.Query("cinemaId"))
	filmId, _ := strconv.Atoi(c.Query("filmId"))
	dayNum, _ := strconv.Atoi(c.Query("dayNum"))

	grpcReq := &film.GetFilmsByCidADayReq{
		CinemaId: int64(cinemaId),
		FilmId:   int64(filmId),
		DayNum:   int64(dayNum),
	}
	grpcRsp := &film.GetFilmsByCidADayRsp{}

	req := client.NewRequest(serviceFilm, endpointGetFilmsByCidADay, grpcReq)

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

func Search(c *gin.Context) {
	grpcReq := &film.SearchReq{}
	grpcRsp := &film.SearchRep{}

	req := client.NewRequest(serviceFilm, endpointSearch, grpcReq)

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
