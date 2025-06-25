package logic

import (
	"context"

	"go-zero-gorm/service/user/rpc/internal/svc"
	"go-zero-gorm/service/user/rpc/types/user"

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
	one, err := l.svcCtx.UserModel.FindOne(context.Background(), in.Id)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoReply{
		Id:     one.Id,
		Name:   one.Name.String,
		CreatedAt: one.CreatedAt.Time.Format("2006-01-02 15:04:05"),
	}, nil
}
