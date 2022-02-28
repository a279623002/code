package server

import (
	cinema "cinema-srv/proto"
	"config"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	user "user-srv/proto"
)

type Grpc struct {
	Service  string
	Endpoint string
	Req      interface{}
	Rsp      interface{}
}

func GetGrpc(c *gin.Context) (grpc *Grpc, err error) {
	path := c.Request.URL.Path
	paths := strings.Split(path, "/")
	if len(paths) != 3 {
		return &Grpc{}, errors.New("bad request")
	}
	if paths[1] == "user" {
		if paths[2] == "registAccount" {
			grpc = RegistAccount(c, config.Namespace+config.ServiceNameUser, "User.RegistAccount")
		}
		if paths[2] == "loginAccount" {
			grpc = LoginAccount(c, config.Namespace+config.ServiceNameUser, "User.LoginAccount")
		}
		if paths[2] == "wantScore" {
			grpc = WantScore(c, config.Namespace+config.ServiceNameUser, "User.WantScore")
		}
		if paths[2] == "updateUserProfile" {
			grpc = UpdateUserProfile(c, config.Namespace+config.ServiceNameUser, "User.UpdateUserProfile")
		}
	}
	if paths[1] == "cinema" {
		if paths[2] == "locationCinema" {
			grpc = LocationCinema(c, config.Namespace+config.ServiceNameCinema, "Cinema.LocationCinema")
		}
		if paths[2] == "getCinemaMessageByCid" {
			grpc = GetCinemaMessageByCid(c, config.Namespace+config.ServiceNameCinema, "Cinema.GetCinemaMessageByCid")
		}
		if paths[2] == "getMovieHallByMHId" {
			grpc = GetMovieHallByMHId(c, config.Namespace+config.ServiceNameCinema, "Cinema.GetMovieHallByMHId")
		}
	}
	if grpc == nil {
		err = errors.New("couldn`t found services")
	}
	return
}

func GetMovieHallByMHId(c *gin.Context, service, endpoint string) *Grpc {
	mhId, _ := strconv.Atoi(c.Query("mhId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &cinema.GetMovieHallByMHIdReq{
			MhId: int64(mhId),
		},
		Rsp: &cinema.GetMovieHallByMHIdRsp{},
	}
}

func GetCinemaMessageByCid(c *gin.Context, service, endpoint string) *Grpc {
	cinemaId, _ := strconv.Atoi(c.Query("cinemaId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &cinema.GetCinemaMessageByCidReq{
			CinemaId: int64(cinemaId),
		},
		Rsp: &cinema.GetCinemaMessageByCidRsp{},
	}
}

func LocationCinema(c *gin.Context, service, endpoint string) *Grpc {
	locationId, _ := strconv.Atoi(c.Query("locationId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &cinema.LocationCinemaReq{
			LocationId: int64(locationId),
		},
		Rsp: &cinema.LocationCinemaRsp{},
	}
}

func UpdateUserProfile(c *gin.Context, service, endpoint string) *Grpc {
	userName := c.Query("userName")
	userEmail := c.Query("userEmail")
	userPhone := c.Query("userPhone")
	userId, _ := strconv.Atoi(c.Query("userId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &user.UpdateUserProfileReq{
			UserName:  userName,
			UserEmail: userEmail,
			UserPhone: userPhone,
			UserID:    int64(userId),
		},
		Rsp: &user.UpdateUserProfileRsp{},
	}
}

func WantScore(c *gin.Context, service, endpoint string) *Grpc {
	userId, _ := strconv.Atoi(c.Query("userId"))
	movieId, _ := strconv.Atoi(c.Query("movieId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &user.WantScoreReq{
			UserId:  int64(userId),
			MovieId: int64(movieId),
		},
		Rsp: &user.WantScoreRsp{},
	}
}
func LoginAccount(c *gin.Context, service, endpoint string) *Grpc {
	email := c.Query("email")
	password := c.Query("password")
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &user.LoginAccountReq{
			Email:    email,
			Password: password,
		},
		Rsp: &user.LoginAccountRsp{},
	}
}

func RegistAccount(c *gin.Context, service, endpoint string) *Grpc {
	email := c.Query("email")
	userName := c.Query("username")
	password := c.Query("password")
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &user.RegistAccountReq{
			Email:    email,
			UserName: userName,
			Password: password,
		},
		Rsp: &user.RegistAccountRsp{},
	}
}
