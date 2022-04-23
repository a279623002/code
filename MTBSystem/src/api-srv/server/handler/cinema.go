package handler

import (
	cinema "cinema-srv/proto"
	"config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"net/http"
	"strconv"
)

var (
	serviceCinema                 = config.Namespace + config.ServiceNameCinema
	endpointGetMovieHallByMHId    = "Cinema.GetMovieHallByMHId"
	endpointGetCinemaMessageByCid = "Cinema.GetCinemaMessageByCid"
	endpointLocationCinema        = "Cinema.LocationCinema"
)

func LocationCinema(c *gin.Context) {
	locationId, _ := strconv.Atoi(c.Query("locationId"))
	grpcReq := &cinema.LocationCinemaReq{
		LocationId: int64(locationId),
	}
	grpcRsp := &cinema.LocationCinemaRsp{}

	req := client.NewRequest(serviceCinema, endpointLocationCinema, grpcReq)

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

func GetCinemaMessageByCid(c *gin.Context) {
	cinemaId, _ := strconv.Atoi(c.Query("cinemaId"))
	grpcReq := &cinema.GetCinemaMessageByCidReq{
		CinemaId: int64(cinemaId),
	}
	grpcRsp := &cinema.GetCinemaMessageByCidRsp{}

	req := client.NewRequest(serviceCinema, endpointGetCinemaMessageByCid, grpcReq)

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

func GetMovieHallByMHId(c *gin.Context) {
	mhId, _ := strconv.Atoi(c.Query("mhId"))
	grpcReq := &cinema.GetMovieHallByMHIdReq{
		MhId: int64(mhId),
	}
	grpcRsp := &cinema.GetMovieHallByMHIdRsp{}

	req := client.NewRequest(serviceCinema, endpointGetMovieHallByMHId, grpcReq)

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
