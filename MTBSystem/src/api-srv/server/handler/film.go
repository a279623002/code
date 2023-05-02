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
	serviceFilm                   = config.Namespace + config.ServiceNameFilm
	endpointSearch                = "Film.Search"
	endpointGetFilmsByCidADay     = "Film.GetFilmsByCidADay"
	endpointMovieComingNew        = "Film.MovieComingNew"
	endpointLocationMovies        = "Film.LocationMovies"
	endpointImageAll              = "Film.ImageAll"
	endpointMovieCreditsWithTypes = "Film.MovieCreditsWithTypes"
	endpointMovieDetail           = "Film.MovieDetail"
	endpointHotPlayMovies         = "Film.HotPlayMovies"
)

// @Summary 正在售票列表
// @Tags 影片中心
// @Description 正在售票列表
// @Accept json
// @Produce json
// @Success 1 {object} model.Response "success"
// @Router /film/hotPlayMovies [get]
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

// @Summary 影片详情
// @Tags 影片中心
// @Description 影片详情
// @Accept json
// @Produce json
// @Param movieId query int true "影片id"
// @Success 1 {object} model.Response "success"
// @Router /film/movieDetail [post]
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

// @Summary 获取影片导演演员
// @Tags 影片中心
// @Description 获取影片导演演员
// @Accept json
// @Produce json
// @Param movieId query int true "影片id"
// @Success 1 {object} model.Response "success"
// @Router /film/movieCreditsWithTypes [post]
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

// @Summary 剧照
// @Tags 影片中心
// @Description 剧照
// @Accept json
// @Produce json
// @Param movieId query int true "影片id"
// @Success 1 {object} model.Response "success"
// @Router /film/imageAll [post]
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

// @Summary 上映的影片
// @Tags 影片中心
// @Description 上映的影片
// @Accept json
// @Produce json
// @Success 1 {object} model.Response "success"
// @Router /film/locationMovies [get]
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

// @Summary 即将上映的影片
// @Tags 影片中心
// @Description 即将上映的影片
// @Accept json
// @Produce json
// @Success 1 {object} model.Response "success"
// @Router /film/movieComingNew [post]
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

// @Summary 获取正在销售的影片信息
// @Tags 影片中心
// @Description 获取正在销售的影片信息
// @Accept json
// @Produce json
// @Param cinemaId query int true "影厅id"
// @Param filmId query int true "影片id"
// @Param dayNum query int true "时间 0：今天 1：明天 2：后天"
// @Success 1 {object} model.Response "success"
// @Router /film/getFilmsByCidADay [post]
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

// @Summary 搜索
// @Tags 影片中心
// @Description 搜索--暂时写死
// @Accept json
// @Produce json
// @Success 1 {object} model.Response "success"
// @Router /film/hotPlayMovies [get]
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
