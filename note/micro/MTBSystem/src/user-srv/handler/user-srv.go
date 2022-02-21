package handler

import (
	"context"
	pb "user-srv/proto"
)

type UserHandler struct {

}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (c *UserHandler) InsertUser(ctx context.Context, req *pb.InsertUserReq, resp *pb.InsertUserResp) error {
	return nil
}

func (c *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserReq, resp *pb.DeleteUserResp) error {
	return nil
}

func (c *UserHandler) SelectUser(ctx context.Context, req *pb.SelectUserReq, resp *pb.SelectUserResp) error {
	resp = &pb.SelectUserResp{User:&pb.User{Id:1}}
	return nil
}

func (c *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserReq, resp *pb.UpdateUserResp) error {
	return nil
}


var _ pb.UserSrvHandler = (*UserHandler)(nil)