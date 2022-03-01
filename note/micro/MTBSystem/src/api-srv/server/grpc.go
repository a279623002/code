package server

import (
	cinema "cinema-srv/proto"
	"config"
	"errors"
	film "film-srv/proto"
	"github.com/gin-gonic/gin"
	place "place-srv/proto"
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
	if paths[1] == "place" {
		if paths[2] == "hotCitiesByCinema" {
			grpc = HotCitiesByCinema(c, config.Namespace+config.ServiceNamePlace, "Place.HotCitiesByCinema")
		}
	}

	if paths[1] == "film" {
		if paths[2] == "hotPlayMovies" {
			grpc = HotPlayMovies(c, config.Namespace+config.ServiceNameFilm, "Film.HotPlayMovies")
		}
		if paths[2] == "movieDetail" {
			grpc = MovieDetail(c, config.Namespace+config.ServiceNameFilm, "Film.MovieDetail")
		}
		if paths[2] == "movieCreditsWithTypes" {
			grpc = MovieCreditsWithTypes(c, config.Namespace+config.ServiceNameFilm, "Film.MovieCreditsWithTypes")
		}
		if paths[2] == "imageAll" {
			grpc = ImageAll(c, config.Namespace+config.ServiceNameFilm, "Film.ImageAll")
		}
		if paths[2] == "locationMovies" {
			grpc = LocationMovies(c, config.Namespace+config.ServiceNameFilm, "Film.LocationMovies")
		}
		if paths[2] == "movieComingNew" {
			grpc = MovieComingNew(c, config.Namespace+config.ServiceNameFilm, "Film.MovieComingNew")
		}
		if paths[2] == "getFilmsByCidADay" {
			grpc = GetFilmsByCidADay(c, config.Namespace+config.ServiceNameFilm, "Film.GetFilmsByCidADay")
		}
		if paths[2] == "search" {
			grpc = Search(c, config.Namespace+config.ServiceNameFilm, "Film.Search")
		}
	}

	if grpc == nil {
		err = errors.New("couldn`t found services")
	}
	return
}


func Search(c *gin.Context, service, endpoint string) *Grpc {
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &film.SearchReq{},
		Rsp:      &film.SearchRep{},
	}
}

func GetFilmsByCidADay(c *gin.Context, service, endpoint string) *Grpc {
	cinemaId, _ := strconv.Atoi(c.Query("cinemaId"))
	filmId, _ := strconv.Atoi(c.Query("filmId"))
	dayNum, _ := strconv.Atoi(c.Query("dayNum"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &film.GetFilmsByCidADayReq{
			CinemaId: int64(cinemaId),
			FilmId:   int64(filmId),
			DayNum:   int64(dayNum),
		},
		Rsp: &film.GetFilmsByCidADayRsp{},
	}
}

func MovieComingNew(c *gin.Context, service, endpoint string) *Grpc {
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &film.MovieComingNewReq{},
		Rsp:      &film.MovieComingNewRep{},
	}
}

func LocationMovies(c *gin.Context, service, endpoint string) *Grpc {
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &film.LocationMoviesReq{},
		Rsp:      &film.LocationMoviesRep{},
	}
}

func ImageAll(c *gin.Context, service, endpoint string) *Grpc {
	movieId, _ := strconv.Atoi(c.Query("movieId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &film.ImageAllReq{
			MovieId: int64(movieId),
		},
		Rsp: &film.ImageAllRep{},
	}
}

func MovieCreditsWithTypes(c *gin.Context, service, endpoint string) *Grpc {
	movieId, _ := strconv.Atoi(c.Query("movieId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &film.MovieCreditsWithTypesReq{
			MovieId: int64(movieId),
		},
		Rsp: &film.MovieCreditsWithTypesRep{},
	}
}

func MovieDetail(c *gin.Context, service, endpoint string) *Grpc {
	movieId, _ := strconv.Atoi(c.Query("movieId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req: &film.MovieDetailReq{
			MovieId: int64(movieId),
		},
		Rsp: &film.MovieDetailRep{},
	}
}

func HotPlayMovies(c *gin.Context, service, endpoint string) *Grpc {
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &film.HotPlayMoviesReq{},
		Rsp:      &film.HotPlayMoviesRep{},
	}
}

func HotCitiesByCinema(c *gin.Context, service, endpoint string) *Grpc {
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &place.HotCitiesByCinemaReq{},
		Rsp:      &place.HotCitiesByCinemaRep{},
	}
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
