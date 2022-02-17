package handler

import (
	"context"

	pb "usersrv/proto"
)

type UserHandler struct {

}
// new一个UserHandler
func NewUserHandler() *UserHandler{
	return &UserHandler{}
}

func (c *UserHandler) InsertUser(ctx context.Context, req * pb.InsertUserReq,rsp *pb.InsertUserRep)error {
	return nil
}

func (c *UserHandler) DeletetUser(ctx context.Context, req * pb.DeletetUserReq,rsp *pb.DeletetUserRep)error {
	return nil
}

func (c *UserHandler) SelectUser(ctx context.Context, req * pb.SelectUserReq,rsp *pb.SelectUserRep)error {
	return nil
}

func (c *UserHandler) UpdateUser(ctx context.Context, req * pb.UpdateUserReq,rsp *pb.UpdateUserRep)error {
	return nil
}
