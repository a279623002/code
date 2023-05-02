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

// @Summary 获取影院
// @Tags 影厅中心
// @Description 根据location_id 获取影院
// @Accept json
// @Produce json
// @Param locationId query int true "定位id"
// @Success 1 {object} model.Response "success"
// @Router /cinema/locationCinema [post]
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

// @Summary 正在销售的影片信息和影院信息
// @Tags 影厅中心
// @Description 正在销售的影片信息和影院信息
// @Accept json
// @Produce json
// @Param cinemaId query int true "影院id"
// @Success 1 {object} model.Response "success"
// @Router /cinema/getCinemaMessageByCid [post]
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

// @Summary 获取影厅座位表
// @Tags 影厅中心
// @Description 获取影厅座位表
// @Accept json
// @Produce json
// @Param mhId query int true "影厅id"
// @Success 1 {object} model.Response "success"
// @Router /cinema/getMovieHallByMHId [post]
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
