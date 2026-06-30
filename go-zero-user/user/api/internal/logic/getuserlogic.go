package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-user/user/api/internal/svc"
	"go-zero-user/user/api/internal/types"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (*types.GetUserResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.GetUserResp{
		User: types.User{
			Id:        user.Id,
			Username:  user.Username,
			Phone:     user.Phone,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
