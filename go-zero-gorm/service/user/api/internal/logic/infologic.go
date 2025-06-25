package logic

import (
	"context"

	"go-zero-gorm/service/user/api/internal/svc"
	"go-zero-gorm/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	user "go-zero-gorm/service/user/rpc/userclient"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.InfoReq) (resp *types.InfoReply, err error) {
	info, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{Id: req.Id})
	if err != nil {
		return
	}
	resp = &types.InfoReply{
		Id:         info.Id,
		Name:       info.Name,
		Created_at: info.CreatedAt,
	}
	return
}
