package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-user/user/api/internal/model"
	"go-zero-user/user/api/internal/svc"
	"go-zero-user/user/api/internal/types"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (*types.CreateUserResp, error) {
	if req.Username == "" || req.Phone == "" {
		return nil, fmt.Errorf("username and phone are required")
	}

	user := &model.User{
		Username: req.Username,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := l.svcCtx.UserModel.Create(l.ctx, user); err != nil {
		return nil, err
	}

	return &types.CreateUserResp{Id: user.Id}, nil
}
