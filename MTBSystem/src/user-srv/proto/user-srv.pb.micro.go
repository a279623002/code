// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user-srv.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	// 注册用户
	RegistAccount(ctx context.Context, in *RegistAccountReq, opts ...client.CallOption) (*RegistAccountRsp, error)
	// 用户登录
	LoginAccount(ctx context.Context, in *LoginAccountReq, opts ...client.CallOption) (*LoginAccountRsp, error)
	// 密码重置
	ResetAccount(ctx context.Context, in *ResetAccountReq, opts ...client.CallOption) (*ResetAccountRsp, error)
	// 评分
	WantScore(ctx context.Context, in *WantScoreReq, opts ...client.CallOption) (*WantScoreRsp, error)
	// 修改用户信息
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...client.CallOption) (*UpdateUserProfileRsp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) RegistAccount(ctx context.Context, in *RegistAccountReq, opts ...client.CallOption) (*RegistAccountRsp, error) {
	req := c.c.NewRequest(c.name, "User.RegistAccount", in)
	out := new(RegistAccountRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) LoginAccount(ctx context.Context, in *LoginAccountReq, opts ...client.CallOption) (*LoginAccountRsp, error) {
	req := c.c.NewRequest(c.name, "User.LoginAccount", in)
	out := new(LoginAccountRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ResetAccount(ctx context.Context, in *ResetAccountReq, opts ...client.CallOption) (*ResetAccountRsp, error) {
	req := c.c.NewRequest(c.name, "User.ResetAccount", in)
	out := new(ResetAccountRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) WantScore(ctx context.Context, in *WantScoreReq, opts ...client.CallOption) (*WantScoreRsp, error) {
	req := c.c.NewRequest(c.name, "User.WantScore", in)
	out := new(WantScoreRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...client.CallOption) (*UpdateUserProfileRsp, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUserProfile", in)
	out := new(UpdateUserProfileRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	// 注册用户
	RegistAccount(context.Context, *RegistAccountReq, *RegistAccountRsp) error
	// 用户登录
	LoginAccount(context.Context, *LoginAccountReq, *LoginAccountRsp) error
	// 密码重置
	ResetAccount(context.Context, *ResetAccountReq, *ResetAccountRsp) error
	// 评分
	WantScore(context.Context, *WantScoreReq, *WantScoreRsp) error
	// 修改用户信息
	UpdateUserProfile(context.Context, *UpdateUserProfileReq, *UpdateUserProfileRsp) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		RegistAccount(ctx context.Context, in *RegistAccountReq, out *RegistAccountRsp) error
		LoginAccount(ctx context.Context, in *LoginAccountReq, out *LoginAccountRsp) error
		ResetAccount(ctx context.Context, in *ResetAccountReq, out *ResetAccountRsp) error
		WantScore(ctx context.Context, in *WantScoreReq, out *WantScoreRsp) error
		UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, out *UpdateUserProfileRsp) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) RegistAccount(ctx context.Context, in *RegistAccountReq, out *RegistAccountRsp) error {
	return h.UserHandler.RegistAccount(ctx, in, out)
}

func (h *userHandler) LoginAccount(ctx context.Context, in *LoginAccountReq, out *LoginAccountRsp) error {
	return h.UserHandler.LoginAccount(ctx, in, out)
}

func (h *userHandler) ResetAccount(ctx context.Context, in *ResetAccountReq, out *ResetAccountRsp) error {
	return h.UserHandler.ResetAccount(ctx, in, out)
}

func (h *userHandler) WantScore(ctx context.Context, in *WantScoreReq, out *WantScoreRsp) error {
	return h.UserHandler.WantScore(ctx, in, out)
}

func (h *userHandler) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, out *UpdateUserProfileRsp) error {
	return h.UserHandler.UpdateUserProfile(ctx, in, out)
}