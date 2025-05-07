package logic

import (
	"context"

	"go-zero-shiro/service/user/rpc/internal/svc"
	"go-zero-shiro/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.UserModel.FindOne(in.Id)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoReply{
		Id:     one.ID,
		Name:   one.Name,
		// Number: one.Number,
		// Gender: one.Gender,
	}, nil
}
