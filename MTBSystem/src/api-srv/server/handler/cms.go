package handler

import (
	cms "cms-srv/proto"
	"config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"net/http"
	"strconv"
)

var (
	serviceCms               = config.Namespace + config.ServiceNameCMS
	endpointUserLogin        = "Cms.UserLogin"
	endpointUpdateMessage    = "Cms.UpdateMessage"
	endpointAllFilms         = "Cms.AllFilms"
	endpointAllUsers         = "Cms.AllUsers"
	endpointAllAdminUsers    = "Cms.AllAdminUsers"
	endpointAllComments      = "Cms.AllComments"
	endpointAllOrders        = "Cms.AllOrders"
	endpointAllAddress       = "Cms.AllAddress"
	endpointAddFilm          = "Cms.AddFilm"
	endpointUpdateFilm       = "Cms.UpdateFilm"
	endpointDeleteFilm       = "Cms.DeleteFilm"
	endpointAddAdminUser     = "Cms.AddAdminUser"
	endpointAddAddress       = "Cms.AddAddress"
	endpointUpdateAddress    = "Cms.UpdateAddress"
	endpointDeleteAddress    = "Cms.DeleteAddress"
	endpointDeleteAdminUser  = "Cms.DeleteAdminUser"
	endpointAllMovieHall     = "Cms.AllMovieHall"
	endpointAddMovieHall     = "Cms.AddMovieHall"
	endpointUpdateMovieHall  = "Cms.UpdateMovieHall"
	endpointDeleteMovieHall  = "Cms.DeleteMovieHall"
	endpointAllCinemaFilms   = "Cms.AllCinemaFilms"
	endpointAddCinemaFilm    = "Cms.AddCinemaFilm"
	endpointUpdateCinemaFilm = "Cms.UpdateCinemaFilm"
	endpointDeleteCinemaFilm = "Cms.DeleteCinemaFilm"
	endpointRegisterCinema   = "Cms.RegisterCinema"
	endpointAllCinemaHall    = "Cms.AllCinemaHall"
)

func AllCinemaHall(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	cinemaID, _ := strconv.Atoi(c.Query("cinemaID"))
	grpcReq := &cms.AllCinemaHallReq{
		CinemaID: int64(cinemaID),
		AdminID:  int64(adminID),
	}
	grpcRsp := &cms.AllCinemaHallRsp{}

	req := client.NewRequest(serviceCms, endpointAllCinemaHall, grpcReq)

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

func RegisterCinema(c *gin.Context) {
	cinemaName := c.Query("cinemaName")
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	cinemaAddress := c.Query("cinemaAddress")
	locationID, _ := strconv.Atoi(c.Query("locationID"))
	cinemaTypes := c.Query("cinemaTypes")
	cinemaCard, _ := strconv.Atoi(c.Query("cinemaCard"))
	cinemaMinPrice, _ := strconv.Atoi(c.Query("cinemaMinPrice"))
	cinemaSupport := c.Query("cinemaSupport")
	cinemaDiscount, _ := strconv.Atoi(c.Query("cinemaDiscount"))
	cinemaPhone, _ := strconv.Atoi(c.Query("cinemaPhone"))
	grpcReq := &cms.RegisterCinemaReq{
		AdminID:        int64(adminID),
		CinemaName:     cinemaName,
		CinemaAddress:  cinemaAddress,
		LocationID:     int64(locationID),
		CinemaTypes:    cinemaTypes,
		CinemaCard:     int64(cinemaCard),
		CinemaMinPrice: int64(cinemaMinPrice),
		CinemaSupport:  cinemaSupport,
		CinemaDiscount: int64(cinemaDiscount),
		CinemaPhone:    int64(cinemaPhone),
	}
	grpcRsp := &cms.RegisterCinemaRsp{}

	req := client.NewRequest(serviceCms, endpointRegisterCinema, grpcReq)

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

func DeleteCinemaFilm(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	cfId, _ := strconv.Atoi(c.Query("cfId"))
	grpcReq := &cms.DeleteCinemaFilmReq{
		AdminID: int64(adminID),
		CfId:    int64(cfId),
	}
	grpcRsp := &cms.DeleteCinemaFilmRsp{}

	req := client.NewRequest(serviceCms, endpointDeleteCinemaFilm, grpcReq)

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

func UpdateCinemaFilm(c *gin.Context) {
	cinemaID, _ := strconv.Atoi(c.Query("cinemaID"))
	filmID, _ := strconv.Atoi(c.Query("filmID"))
	hallID, _ := strconv.Atoi(c.Query("hallID"))
	filmName := c.Query("filmName")
	cinemaName := c.Query("cinemaName")
	releaseTimeYear, _ := strconv.Atoi(c.Query("releaseTimeYear"))
	releaseTimeMonth, _ := strconv.Atoi(c.Query("releaseTimeMonth"))
	releaseTimeDay, _ := strconv.Atoi(c.Query("releaseTimeDay"))
	releaseTime := c.Query("releaseTime")
	t := c.Query("type")
	releaseAdd := c.Query("releaseAdd")
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	length, _ := strconv.Atoi(c.Query("length"))
	releaseDiscount, _ := strconv.ParseFloat(c.Query("releaseDiscount"), 32)
	cfId, _ := strconv.Atoi(c.Query("cfId"))
	grpcReq := &cms.UpdateCinemaFilmReq{
		CinemaID:         int64(cinemaID),
		FilmID:           int64(filmID),
		HallID:           int64(hallID),
		FilmName:         filmName,
		CinemaName:       cinemaName,
		ReleaseTimeYear:  int64(releaseTimeYear),
		ReleaseTimeMonth: int64(releaseTimeMonth),
		ReleaseTimeDay:   int64(releaseTimeDay),
		ReleaseTime:      releaseTime,
		ReleaseType:      t,
		ReleaseAdd:       releaseAdd,
		AdminID:          int64(adminID),
		Length:           int64(length),
		ReleaseDiscount:  float32(releaseDiscount),
		CfID:             int64(cfId),
	}
	grpcRsp := &cms.UpdateCinemaFilmRsp{}

	req := client.NewRequest(serviceCms, endpointUpdateCinemaFilm, grpcReq)

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

func AddCinemaFilm(c *gin.Context) {
	cinemaID, _ := strconv.Atoi(c.Query("cinemaID"))
	movieID, _ := strconv.Atoi(c.Query("movieID"))
	hallID, _ := strconv.Atoi(c.Query("hallID"))
	titleCn := c.Query("titleCn")
	cinemaName := c.Query("cinemaName")
	releaseTimeYear, _ := strconv.Atoi(c.Query("releaseTimeYear"))
	releaseTimeMonth, _ := strconv.Atoi(c.Query("releaseTimeMonth"))
	releaseTimeDay, _ := strconv.Atoi(c.Query("releaseTimeDay"))
	releaseTime := c.Query("releaseTime")
	t := c.Query("type")
	releaseAdd := c.Query("releaseAdd")
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	length, _ := strconv.Atoi(c.Query("length"))
	releaseDiscount, _ := strconv.ParseFloat(c.Query("releaseDiscount"), 32)
	grpcReq := &cms.AddCinemaFilmReq{
		CinemaID:         int64(cinemaID),
		MovieID:          int64(movieID),
		HallID:           int64(hallID),
		TitleCn:          titleCn,
		CinemaName:       cinemaName,
		ReleaseTimeYear:  int64(releaseTimeYear),
		ReleaseTimeMonth: int64(releaseTimeMonth),
		ReleaseTimeDay:   int64(releaseTimeDay),
		ReleaseTime:      releaseTime,
		Type:             t,
		ReleaseAdd:       releaseAdd,
		AdminID:          int64(adminID),
		Length:           int64(length),
		ReleaseDiscount:  float32(releaseDiscount),
	}
	grpcRsp := &cms.AddCinemaFilmRsp{}

	req := client.NewRequest(serviceCms, endpointAddCinemaFilm, grpcReq)

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

func AllCinemaFilms(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	page, _ := strconv.Atoi(c.Query("page"))
	grpcReq := &cms.AllCinemaFilmsReq{
		Page:    int64(page),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.AllCinemaFilmsRsp{}

	req := client.NewRequest(serviceCms, endpointAllCinemaFilms, grpcReq)

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

func DeleteMovieHall(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	mhId, _ := strconv.Atoi(c.Query("mhId"))
	grpcReq := &cms.DeleteMovieHallReq{
		AdminID: int64(adminID),
		MhId:    int64(mhId),
	}
	grpcRsp := &cms.DeleteMovieHallRsp{}

	req := client.NewRequest(serviceCms, endpointDeleteMovieHall, grpcReq)

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

func UpdateMovieHall(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	cinemaId, _ := strconv.Atoi(c.Query("cinemaId"))
	mhName := c.Query("mhName")
	mhAddress := c.Query("mhAddress")
	grpcReq := &cms.UpdateMovieHallReq{
		AdminID:   int64(adminID),
		MhName:    mhName,
		MhAddress: mhAddress,
		CinemaId:  int64(cinemaId),
	}
	grpcRsp := &cms.UpdateMovieHallRsp{}

	req := client.NewRequest(serviceCms, endpointUpdateMovieHall, grpcReq)

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

func AddMovieHall(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	cinemaId, _ := strconv.Atoi(c.Query("cinemaId"))
	mhName := c.Query("mhName")
	mhAddress := c.Query("mhAddress")
	grpcReq := &cms.AddMovieHallReq{
		AdminID:   int64(adminID),
		MhName:    mhName,
		MhAddress: mhAddress,
		CinemaId:  int64(cinemaId),
	}
	grpcRsp := &cms.AddMovieHallRsp{}

	req := client.NewRequest(serviceCms, endpointAddMovieHall, grpcReq)

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

func AllMovieHall(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	page, _ := strconv.Atoi(c.Query("page"))
	grpcReq := &cms.AllMovieHallReq{
		Page:    int64(page),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.AllMovieHallRsp{}

	req := client.NewRequest(serviceCms, endpointAllMovieHall, grpcReq)

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

func DeleteAdminUser(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	auID, _ := strconv.Atoi(c.Query("auID"))
	grpcReq := &cms.DeleteAdminUserReq{
		AuID:    int64(auID),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.DeleteAdminUserRsp{}

	req := client.NewRequest(serviceCms, endpointDeleteAdminUser, grpcReq)

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

func DeleteAddress(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	id, _ := strconv.Atoi(c.Query("id"))
	grpcReq := &cms.DeleteAddressReq{
		Id:      int64(id),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.DeleteAddressRsp{}

	req := client.NewRequest(serviceCms, endpointDeleteAddress, grpcReq)

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

func UpdateAddress(c *gin.Context) {
	name := c.Query("name")
	pinyinFull := c.Query("pinyinFull")
	pinyinShort := c.Query("pinyinShort")
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	id, _ := strconv.Atoi(c.Query("id"))
	grpcReq := &cms.UpdateAddressReq{
		Id:          int64(id),
		Name:        name,
		PinyinFull:  pinyinFull,
		PinyinShort: pinyinShort,
		AdminID:     int64(adminID),
	}
	grpcRsp := &cms.UpdateAddressRsp{}

	req := client.NewRequest(serviceCms, endpointUpdateAddress, grpcReq)

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

func AddAddress(c *gin.Context) {
	name := c.Query("name")
	pinyinFull := c.Query("pinyinFull")
	pinyinShort := c.Query("pinyinShort")
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	grpcReq := &cms.AddAddressReq{
		AdminID:     int64(adminID),
		Name:        name,
		PinyinFull:  pinyinFull,
		PinyinShort: pinyinShort,
	}
	grpcRsp := &cms.AddAddressRsp{}

	req := client.NewRequest(serviceCms, endpointAddAddress, grpcReq)

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

func AddAdminUser(c *gin.Context) {
	adminName := c.Query("adminName")
	adminPassword := c.Query("adminPassword")
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	adminCinemaID, _ := strconv.Atoi(c.Query("adminCinemaID"))
	adminNum, _ := strconv.Atoi(c.Query("adminNum"))
	grpcReq := &cms.AddAdminUserReq{
		AdminID:       int64(adminID),
		AdminName:     adminName,
		AdminPassword: adminPassword,
		AdminCinemaID: int64(adminCinemaID),
		AdminNum:      int64(adminNum),
	}
	grpcRsp := &cms.AddAdminUserRsp{}

	req := client.NewRequest(serviceCms, endpointAddAdminUser, grpcReq)

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

func DeleteFilm(c *gin.Context) {
	movieID, _ := strconv.Atoi(c.Query("movieID"))
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	grpcReq := &cms.DeleteFilmReq{
		MovieID: int64(movieID),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.DeleteFilmRsp{}

	req := client.NewRequest(serviceCms, endpointDeleteFilm, grpcReq)

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

func UpdateFilm(c *gin.Context) {
	movieID, _ := strconv.Atoi(c.Query("movieID"))
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	img := c.Query("img")
	length, _ := strconv.Atoi(c.Query("length"))
	filmPrice, _ := strconv.ParseFloat(c.Query("filmPrice"), 32)
	filmDirector := c.Query("filmDirector")
	titleCn := c.Query("titleCn")
	titleEn := c.Query("titleEn")
	t := c.Query("type")
	filmDrama := c.Query("filmDrama")
	commonSpecial := c.Query("commonSpecial")
	companyIssued := c.Query("companyIssued")
	country := c.Query("country")
	isTicking, _ := strconv.Atoi(c.Query("isTicking"))
	rDay, _ := strconv.Atoi(c.Query("rDay"))
	rMonth, _ := strconv.Atoi(c.Query("rMonth"))
	rYear, _ := strconv.Atoi(c.Query("rYear"))
	grpcReq := &cms.UpdateFilmReq{
		MovieID:       int64(movieID),
		Img:           img,
		Length:        int64(length),
		FilmPrice:     float32(filmPrice),
		FilmDirector:  filmDirector,
		TitleCn:       titleCn,
		TitleEn:       titleEn,
		Type:          t,
		FilmDrama:     filmDrama,
		CommonSpecial: commonSpecial,
		CompanyIssued: companyIssued,
		Country:       country,
		RDay:          int64(rDay),
		RMonth:        int64(rMonth),
		RYear:         int64(rYear),
		AdminID:       int64(adminID),
		IsTicking:     int64(isTicking),
	}
	grpcRsp := &cms.UpdateFilmRsp{}

	req := client.NewRequest(serviceCms, endpointUpdateFilm, grpcReq)

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

func AddFilm(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	img := c.Query("img")
	length, _ := strconv.Atoi(c.Query("length"))
	filmPrice, _ := strconv.ParseFloat(c.Query("filmPrice"), 32)
	filmScreenwriter := c.Query("filmScreenwriter")
	filmDirector := c.Query("filmDirector")
	titleCn := c.Query("titleCn")
	titleEn := c.Query("titleEn")
	t := c.Query("type")
	filmDrama := c.Query("filmDrama")
	commonSpecial := c.Query("commonSpecial")
	companyIssued := c.Query("companyIssued")
	country := c.Query("country")
	is3D, _ := strconv.Atoi(c.Query("is3D"))
	isDMAX, _ := strconv.Atoi(c.Query("isDMAX"))
	isIMAX, _ := strconv.Atoi(c.Query("isIMAX"))
	isIMAX3D, _ := strconv.Atoi(c.Query("isIMAX3D"))
	rDay, _ := strconv.Atoi(c.Query("rDay"))
	rMonth, _ := strconv.Atoi(c.Query("rMonth"))
	rYear, _ := strconv.Atoi(c.Query("rYear"))
	filmDirectorImg := c.Query("filmDirectorImg")
	filmActor1 := c.Query("filmActor1")
	filmActor1Img := c.Query("filmActor1Img")
	filmActor2 := c.Query("filmActor2")
	filmActor2Img := c.Query("filmActor2Img")
	grpcReq := &cms.AddFilmReq{
		AdminID:          int64(adminID),
		Img:              img,
		Length:           int64(length),
		FilmPrice:        float32(filmPrice),
		FilmScreenwriter: filmScreenwriter,
		FilmDirector:     filmDirector,
		TitleCn:          titleCn,
		TitleEn:          titleEn,
		Type:             t,
		FilmDrama:        filmDrama,
		CommonSpecial:    commonSpecial,
		CompanyIssued:    companyIssued,
		Country:          country,
		Is3D:             int64(is3D),
		IsDMAX:           int64(isDMAX),
		IsIMAX:           int64(isIMAX),
		IsIMAX3D:         int64(isIMAX3D),
		RDay:             int64(rDay),
		RMonth:           int64(rMonth),
		RYear:            int64(rYear),
		FilmDirectorImg:  filmDirectorImg,
		FilmActor1:       filmActor1,
		FilmActor1Img:    filmActor1Img,
		FilmActor2:       filmActor2,
		FilmActor2Img:    filmActor2Img,
	}
	grpcRsp := &cms.AddFilmRsp{}

	req := client.NewRequest(serviceCms, endpointAddFilm, grpcReq)

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

func AllAddress(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	grpcReq := &cms.AllAddressReq{
		Page:    int64(page),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.AllAddressRsp{}

	req := client.NewRequest(serviceCms, endpointAllAddress, grpcReq)

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

func AllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	grpcReq := &cms.AllOrdersReq{
		Page:    int64(page),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.AllOrdersRsp{}

	req := client.NewRequest(serviceCms, endpointAllOrders, grpcReq)

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

func AllComments(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	grpcReq := &cms.AllCommentsReq{
		Page:    int64(page),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.AllCommentsRsp{}

	req := client.NewRequest(serviceCms, endpointAllComments, grpcReq)

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

func AllAdminUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	grpcReq := &cms.AllAdminUsersReq{
		Page:    int64(page),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.AllAdminUsersRsp{}

	req := client.NewRequest(serviceCms, endpointAllAdminUsers, grpcReq)

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

func AllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	grpcReq := &cms.AllUsersReq{
		Page:    int64(page),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.AllUsersRsp{}

	req := client.NewRequest(serviceCms, endpointAllUsers, grpcReq)

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

func AllFilms(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	adminID, _ := strconv.Atoi(c.Query("adminID"))
	grpcReq := &cms.AllFilmsReq{
		Page:    int64(page),
		AdminID: int64(adminID),
	}
	grpcRsp := &cms.AllAddressRsp{}

	req := client.NewRequest(serviceCms, endpointAllFilms, grpcReq)

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

func UpdateMessage(c *gin.Context) {
	grpcReq := &cms.UpdateMessageReq{}
	grpcRsp := &cms.UpdateAddressRsp{}

	req := client.NewRequest(serviceCms, endpointUpdateMessage, grpcReq)

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

func UserLogin(c *gin.Context) {
	userName := c.Query("userName")
	password := c.Query("password")
	grpcReq := &cms.UserLoginReq{
		User:     userName,
		Password: password,
	}
	grpcRsp := &cms.UserLoginRsp{}

	req := client.NewRequest(serviceCms, endpointUserLogin, grpcReq)

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
