package handler

import (
	"config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"net/http"
	place "place-srv/proto"
)

var (
	servicePlace              = config.Namespace + config.ServiceNamePlace
	endpointHotCitiesByCinema = "Place.HotCitiesByCinema"
)

func HotCitiesByCinema(c *gin.Context) {
	grpcReq := &place.HotCitiesByCinemaReq{}
	grpcRsp := &place.HotCitiesByCinemaRep{}

	req := client.NewRequest(servicePlace, endpointHotCitiesByCinema, grpcReq)

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
