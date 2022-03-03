package handler

import (
	"context"
	"errors"
	"user-srv/db"
	pb "user-srv/proto"
)

type UserHandler struct {

}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// 账户注册
func (u *UserHandler) RegistAccount(ctx context.Context, req *pb.RegistAccountReq, rsp *pb.RegistAccountRsp) error {

	userName := req.UserName
	password := req.Password
	email := req.Email
	user, err := db.SelectUserByEmail(email)
	if err != nil {
		//u.logger.Error("error", zap.Error(err))
		//return errors.ErrorUserFailed
		return errors.New("操作异常")
	}
	if user != nil {
		//return errors.ErrorUserAlready
		return errors.New("该邮箱已经被注册过了")
	}
	err = db.InsertUser(userName, password, email)
	if err != nil {
		//u.logger.Error("error", zap.Error(err))
		//return errors.ErrorUserFailed
		return errors.New("操作异常")
	}
	return nil
}

func (u *UserHandler) LoginAccount(ctx context.Context, req *pb.LoginAccountReq, rsp *pb.LoginAccountRsp) error {
	email := req.Email
	password := req.Password
	user, err := db.SelectUserByPasswordName(email, password)
	if err != nil {
		//u.logger.Error("error", zap.Error(err))
		//return errors.ErrorUserFailed
		return errors.New("操作异常")
	}
	if user == nil {
		//return errors.ErrorUserLoginFailed
		return errors.New("密码或者用户名错误")
	}
	rsp.Email = user.Email
	rsp.Phone = user.Phone
	rsp.UserID = user.UserId
	rsp.UserName = user.UserName
	return nil
}

func (u *UserHandler) ResetAccount(ctx context.Context, req *pb.ResetAccountReq, rsp *pb.ResetAccountRsp) error {
	return nil
}

func (u *UserHandler) WantScore(ctx context.Context, req *pb.WantScoreReq, rsp *pb.WantScoreRsp) error {

	orderNum, err := db.SelectOrderByUidMid(req.MovieId, req.UserId)
	if err != nil {
		//u.logger.Error("error", zap.Error(err))
		//return errors.ErrorUserFailed
		return errors.New("操作异常")
	}
	if orderNum == "" {
		//u.logger.Error("error", zap.Error(err))
		//return errors.ErrorScoreForbid
		return errors.New("你没有买过该电影票，无法进行评分")
	}
	err = db.UpdateOrderScore(orderNum, req.Score)
	if err != nil {
		//u.logger.Error("error", zap.Error(err))
		//return errors.ErrorUserFailed
		return errors.New("操作异常")
	}
	return nil
}

func (u *UserHandler) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileReq, rsp *pb.UpdateUserProfileRsp) error {

	userEmail := req.UserEmail
	userName := req.UserName
	userPhone := req.UserPhone
	userID := req.UserID
	if userEmail != "" {
		err := db.UpdateUserEmailProfile(userEmail, userID)
		if err != nil {
			//u.logger.Error("error", zap.Error(err))
			//return errors.ErrorUserFailed
			return errors.New("操作异常")
		}
	}
	if userName != "" {
		err := db.UpdateUserNameProfile(userName, userID)
		if err != nil {
			//u.logger.Error("error", zap.Error(err))
			//return errors.ErrorUserFailed
			return errors.New("操作异常")
		}
	}
	if userPhone != "" {
		err := db.UpdateUserPhoneProfile(userPhone, userID)
		if err != nil {
			//u.logger.Error("error", zap.Error(err))
			//return errors.ErrorUserFailed
			return errors.New("操作异常")
		}
	}
	return nil
}
