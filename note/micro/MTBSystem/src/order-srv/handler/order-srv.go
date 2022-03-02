package handler

import (
	"context"
	"errors"
	"fmt"
	"order-srv/db"
	"order-srv/entity"
	pb "order-srv/proto"
	"strconv"
	"time"
)

type OrderServiceExtHandler struct {

}

func NewOrderServiceExtHandler() *OrderServiceExtHandler {
	return &OrderServiceExtHandler{

	}
}

// 记录想看的信息
func (o *OrderServiceExtHandler) WantTicket(ctx context.Context, req *pb.WantTicketReq, rsp *pb.WantTicketRsp) error {

	err := db.InsertWantSeeRecord(req.FilmId, req.UserId)
	if err != nil {
		return errors.New("操作异常")
	}
	return nil
}

func (o *OrderServiceExtHandler) Ticket(ctx context.Context, req *pb.TicketReq, rsp *pb.TicketRsp) error {

	price, err := db.SelectFilmPrice(req.FilmId)
	if err != nil {
		return errors.New("操作异常")
	}
	orderNum := time.Now().Unix()
	err = db.InsertOrder(strconv.Itoa(int(orderNum)), price, req.MhId, req.UserId, req.FilmId, req.X, req.Y, req.StartTime, req.EndTime)
	if err != nil {
		return errors.New("操作异常")
	}
	rsp.OrderNumD = orderNum
	return nil
}

func (o *OrderServiceExtHandler) PayOrder(ctx context.Context, req *pb.PayOrderReq, rsp *pb.PayOrderRsp) error {

	err := db.UpdateOrderStatus(req.OrderNum, req.UserId)
	if err != nil {
		return errors.New("操作异常")
	}

	err = db.UpdateUserPhone(req.UserId, req.Phone)
	if err != nil {
		return errors.New("操作异常")
	}
	return nil

}

func (o *OrderServiceExtHandler) UndoOrder(ctx context.Context, req *pb.UndoOrderReq, rsp *pb.UndoOrderRsp) error {
	return nil
}

// 查看所有电影票
func (o *OrderServiceExtHandler) LookOrders(ctx context.Context, req *pb.LookOrdersReq, rsp *pb.LookOrdersRsp) error {

	userId := req.UserId
	// 从order.sql获取电影id orderNum mhId startTime
	orders, err := db.SelectOrderNumMovieIdMHIdStartTime(userId)
	if err != nil {
		return errors.New("操作异常")
	}
	movieTicketsPB := []*pb.MovieTicket{}
	for _, order := range orders {
		// 从film_id hall_id获取filmName
		cinemafilm, err := db.SelectFilmNameCinemaName(order.MhId, order.MovieId)
		if err != nil {
			return errors.New("操作异常")
		}
		movieTicketPB := pb.MovieTicket{
			FilmName:  cinemafilm.FilmName,
			Cinema:    cinemafilm.CinemaName,
			StartTime: order.StartTime,
			OrderNum:  order.OrderNum,
		}
		movieTicketsPB = append(movieTicketsPB, &movieTicketPB)
	}

	rsp.MovieTickets = movieTicketsPB
	return nil
}

// 查看所有看过的电影票
func (o *OrderServiceExtHandler) LookAlreadyOrders(ctx context.Context, req *pb.LookAlreadyOrdersReq, rsp *pb.LookAlreadyOrdersRsp) error {

	userId := req.UserId
	orders, err := db.SelectLookAlreadyOrders(userId)
	if err != nil {
		return errors.New("操作异常")
	}
	var oneNoComment int64 = 0
	movies := []*pb.AlreadyMovie{}
	for _, order := range orders {

		actorNames := []string{}
		film, err := db.SelectFilmDetail(order.MovieId)
		if err != nil {
			return errors.New("操作异常")
		}
		actors, err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			return errors.New("操作异常")
		}
		for _, actor := range actors {
			actorNames = append(actorNames, actor.FilmName)
		}
		movie := pb.AlreadyMovie{
			FilmImg:    film.Img,
			FilmName:   film.TitleCn,
			Time:       fmt.Sprintf("%d-%d-%d", film.RYear, film.RMonth, film.RDay),
			Director:   film.FilmDirector,
			ActorNames: actorNames,
			OrderNum:   order.OrderNum,
		}
		movies = append(movies, &movie)
		if order.OrderScore == -1 {
			oneNoComment = oneNoComment + 1
		}
	}
	rsp.Movies = movies
	rsp.OneNoComment = oneNoComment
	rsp.TotalMovie = int64(len(orders))
	return nil
}

// 订单评分
func (o *OrderServiceExtHandler) OrderComment(ctx context.Context, req *pb.OrderCommentReq, rsp *pb.OrderCommentRsp) error {

	userId := req.UserId
	score := req.Score
	orderNum := req.OrderNum
	content := req.CommentContent
	order, err := db.SelectOrderScore(orderNum, userId)
	if err != nil {
		return errors.New("操作异常")
	}
	if order == nil {
		return errors.New("操作异常")
	}
	if order.OrderScore != -1 {
		return errors.New("已经评分了")
	}
	err = db.UpdateOrderScore(score, userId, orderNum)
	if err != nil {
		return errors.New("操作异常")
	}
	err = db.UpdateFilmScore(order.MovieId, score)
	if err != nil {
		return errors.New("操作异常")
	}
	user, err := db.SelectUserNameByUserId(userId)
	if err != nil {
		return errors.New("操作异常")
	}

	comment := entity.Comment{
		FilmId:   order.MovieId,
		Content:  content,
		Title:    strconv.Itoa(int(score)),
		NickName: user.UserName,
		UserId:   userId,
	}
	err = db.InsertComment(&comment)
	if err != nil {
		return errors.New("操作异常")
	}
	return nil
}

// 根据订单编号获取电影票具体信息
func (o *OrderServiceExtHandler) GetOrderMessage(ctx context.Context, req *pb.GetOrderMessageReq, rsp *pb.GetOrderMessageRsp) error {

	orderNum := req.OrderNum
	userId := req.UserId
	// 通过orderNum 从 order 表获取startTime endTime mh_id movie_id order_x order_y create_at order_price
	order, err := db.SelectOrderMessage(orderNum, userId)
	if err != nil {
		return errors.New("操作异常")
	}

	// 通过movie_id 从 film 获取 title_cn img
	film, err := db.SelectFilmMessage(order.MovieId)
	if err != nil {
		return errors.New("操作异常")
	}

	// 通过mh_id 从 movie_hall 获取 mh_name  cinema_id
	hall, err := db.SelectMHNameMHId(order.MhId)
	if err != nil {
		return errors.New("操作异常")
	}

	// 通过cinema_id 从 cinema 获取 cinema_name cinema_add cinema_phone
	cinema, err := db.SelectCinemaByCid(hall.CinemaId)
	if err != nil {
		return errors.New("操作异常")
	}

	// 通过user_id 从 user 获取 phone
	user, err := db.SelectUserPhoneByUserId(userId)
	if err != nil {
		return errors.New("操作异常")
	}

	ticketDetailPB := &pb.TicketDetail{
		FilmName:      film.TitleCn,
		FilmImg:       film.Img,
		StartTime:     order.StartTime,
		EndTime:       order.EndTime,
		CinemaName:    cinema.CinemaName,
		MhName:        hall.MhName,
		Seat:          fmt.Sprintf("%d排%d座", order.OrderX, order.OrderY),
		OrderNum:      orderNum,
		CinemaAddress: cinema.CinemaAdd,
		Price:         order.OrderPrice,
		CreateAt:      order.CreateAt,
		Phone:         user.Phone,
		CinemaPhone:   cinema.CinemaPhone,
	}

	rsp.TicketDetail = ticketDetailPB
	return nil
}


