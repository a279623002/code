package logic

import (
	"context"
	"errors"
	"strings"

	"go-zero-shiro/service/user/api/internal/svc"
	"go-zero-shiro/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	user "go-zero-shiro/service/user/rpc/userclient"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.Username)

	// 使用user rpc
	_, err = l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: int64(userInfo.Id),
	})

	resp = &types.LoginReply{
        Id:           int64(userInfo.Id),
        Name:         userInfo.Name.String,
        // Gender:       userInfo.Gender,
        // AccessToken:  jwtToken,
        // AccessExpire: now + accessExpire,
        // RefreshAfter: now + accessExpire/2,
    }
	return
}
