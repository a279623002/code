// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cms-srv.proto

package cms

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

// Api Endpoints for Cms service

func NewCmsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Cms service

type CmsService interface {
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...client.CallOption) (*UserLoginRsp, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageReq, opts ...client.CallOption) (*UpdateMessageRsp, error)
	AllFilms(ctx context.Context, in *AllFilmsReq, opts ...client.CallOption) (*AllFilmsRsp, error)
	UpdateFilm(ctx context.Context, in *UpdateFilmReq, opts ...client.CallOption) (*UpdateFilmRsp, error)
	DeleteFilm(ctx context.Context, in *DeleteFilmReq, opts ...client.CallOption) (*DeleteFilmRsp, error)
	AllUsers(ctx context.Context, in *AllUsersReq, opts ...client.CallOption) (*AllUsersRsp, error)
	AllAdminUsers(ctx context.Context, in *AllAdminUsersReq, opts ...client.CallOption) (*AllAdminUsersRsp, error)
	AllComments(ctx context.Context, in *AllCommentsReq, opts ...client.CallOption) (*AllCommentsRsp, error)
	AllOrders(ctx context.Context, in *AllOrdersReq, opts ...client.CallOption) (*AllOrdersRsp, error)
	AllAddress(ctx context.Context, in *AllAddressReq, opts ...client.CallOption) (*AllAddressRsp, error)
	AddFilm(ctx context.Context, in *AddFilmReq, opts ...client.CallOption) (*AddFilmRsp, error)
	AddAdminUser(ctx context.Context, in *AddAdminUserReq, opts ...client.CallOption) (*AddAdminUserRsp, error)
	AddAddress(ctx context.Context, in *AddAddressReq, opts ...client.CallOption) (*AddAddressRsp, error)
	UpdateAddress(ctx context.Context, in *UpdateAddressReq, opts ...client.CallOption) (*UpdateAddressRsp, error)
	DeleteAddress(ctx context.Context, in *DeleteAddressReq, opts ...client.CallOption) (*DeleteAddressRsp, error)
	DeleteAdminUser(ctx context.Context, in *DeleteAdminUserReq, opts ...client.CallOption) (*DeleteAdminUserRsp, error)
	AllMovieHall(ctx context.Context, in *AllMovieHallReq, opts ...client.CallOption) (*AllMovieHallRsp, error)
	AddMovieHall(ctx context.Context, in *AddMovieHallReq, opts ...client.CallOption) (*AddMovieHallRsp, error)
	UpdateMovieHall(ctx context.Context, in *UpdateMovieHallReq, opts ...client.CallOption) (*UpdateMovieHallRsp, error)
	DeleteMovieHall(ctx context.Context, in *DeleteMovieHallReq, opts ...client.CallOption) (*DeleteMovieHallRsp, error)
	AllCinemaFilms(ctx context.Context, in *AllCinemaFilmsReq, opts ...client.CallOption) (*AllCinemaFilmsRsp, error)
	AddCinemaFilm(ctx context.Context, in *AddCinemaFilmReq, opts ...client.CallOption) (*AddCinemaFilmRsp, error)
	UpdateCinemaFilm(ctx context.Context, in *UpdateCinemaFilmReq, opts ...client.CallOption) (*UpdateCinemaFilmRsp, error)
	DeleteCinemaFilm(ctx context.Context, in *DeleteCinemaFilmReq, opts ...client.CallOption) (*DeleteCinemaFilmRsp, error)
	RegisterCinema(ctx context.Context, in *RegisterCinemaReq, opts ...client.CallOption) (*RegisterCinemaRsp, error)
	AllCinemaHall(ctx context.Context, in *AllCinemaHallReq, opts ...client.CallOption) (*AllCinemaHallRsp, error)
}

type cmsService struct {
	c    client.Client
	name string
}

func NewCmsService(name string, c client.Client) CmsService {
	return &cmsService{
		c:    c,
		name: name,
	}
}

func (c *cmsService) UserLogin(ctx context.Context, in *UserLoginReq, opts ...client.CallOption) (*UserLoginRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.UserLogin", in)
	out := new(UserLoginRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) UpdateMessage(ctx context.Context, in *UpdateMessageReq, opts ...client.CallOption) (*UpdateMessageRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.UpdateMessage", in)
	out := new(UpdateMessageRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AllFilms(ctx context.Context, in *AllFilmsReq, opts ...client.CallOption) (*AllFilmsRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AllFilms", in)
	out := new(AllFilmsRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) UpdateFilm(ctx context.Context, in *UpdateFilmReq, opts ...client.CallOption) (*UpdateFilmRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.UpdateFilm", in)
	out := new(UpdateFilmRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) DeleteFilm(ctx context.Context, in *DeleteFilmReq, opts ...client.CallOption) (*DeleteFilmRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.DeleteFilm", in)
	out := new(DeleteFilmRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AllUsers(ctx context.Context, in *AllUsersReq, opts ...client.CallOption) (*AllUsersRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AllUsers", in)
	out := new(AllUsersRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AllAdminUsers(ctx context.Context, in *AllAdminUsersReq, opts ...client.CallOption) (*AllAdminUsersRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AllAdminUsers", in)
	out := new(AllAdminUsersRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AllComments(ctx context.Context, in *AllCommentsReq, opts ...client.CallOption) (*AllCommentsRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AllComments", in)
	out := new(AllCommentsRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AllOrders(ctx context.Context, in *AllOrdersReq, opts ...client.CallOption) (*AllOrdersRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AllOrders", in)
	out := new(AllOrdersRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AllAddress(ctx context.Context, in *AllAddressReq, opts ...client.CallOption) (*AllAddressRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AllAddress", in)
	out := new(AllAddressRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AddFilm(ctx context.Context, in *AddFilmReq, opts ...client.CallOption) (*AddFilmRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AddFilm", in)
	out := new(AddFilmRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AddAdminUser(ctx context.Context, in *AddAdminUserReq, opts ...client.CallOption) (*AddAdminUserRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AddAdminUser", in)
	out := new(AddAdminUserRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AddAddress(ctx context.Context, in *AddAddressReq, opts ...client.CallOption) (*AddAddressRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AddAddress", in)
	out := new(AddAddressRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) UpdateAddress(ctx context.Context, in *UpdateAddressReq, opts ...client.CallOption) (*UpdateAddressRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.UpdateAddress", in)
	out := new(UpdateAddressRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) DeleteAddress(ctx context.Context, in *DeleteAddressReq, opts ...client.CallOption) (*DeleteAddressRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.DeleteAddress", in)
	out := new(DeleteAddressRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) DeleteAdminUser(ctx context.Context, in *DeleteAdminUserReq, opts ...client.CallOption) (*DeleteAdminUserRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.DeleteAdminUser", in)
	out := new(DeleteAdminUserRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AllMovieHall(ctx context.Context, in *AllMovieHallReq, opts ...client.CallOption) (*AllMovieHallRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AllMovieHall", in)
	out := new(AllMovieHallRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AddMovieHall(ctx context.Context, in *AddMovieHallReq, opts ...client.CallOption) (*AddMovieHallRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AddMovieHall", in)
	out := new(AddMovieHallRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) UpdateMovieHall(ctx context.Context, in *UpdateMovieHallReq, opts ...client.CallOption) (*UpdateMovieHallRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.UpdateMovieHall", in)
	out := new(UpdateMovieHallRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) DeleteMovieHall(ctx context.Context, in *DeleteMovieHallReq, opts ...client.CallOption) (*DeleteMovieHallRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.DeleteMovieHall", in)
	out := new(DeleteMovieHallRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AllCinemaFilms(ctx context.Context, in *AllCinemaFilmsReq, opts ...client.CallOption) (*AllCinemaFilmsRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AllCinemaFilms", in)
	out := new(AllCinemaFilmsRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AddCinemaFilm(ctx context.Context, in *AddCinemaFilmReq, opts ...client.CallOption) (*AddCinemaFilmRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AddCinemaFilm", in)
	out := new(AddCinemaFilmRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) UpdateCinemaFilm(ctx context.Context, in *UpdateCinemaFilmReq, opts ...client.CallOption) (*UpdateCinemaFilmRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.UpdateCinemaFilm", in)
	out := new(UpdateCinemaFilmRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) DeleteCinemaFilm(ctx context.Context, in *DeleteCinemaFilmReq, opts ...client.CallOption) (*DeleteCinemaFilmRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.DeleteCinemaFilm", in)
	out := new(DeleteCinemaFilmRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) RegisterCinema(ctx context.Context, in *RegisterCinemaReq, opts ...client.CallOption) (*RegisterCinemaRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.RegisterCinema", in)
	out := new(RegisterCinemaRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsService) AllCinemaHall(ctx context.Context, in *AllCinemaHallReq, opts ...client.CallOption) (*AllCinemaHallRsp, error) {
	req := c.c.NewRequest(c.name, "Cms.AllCinemaHall", in)
	out := new(AllCinemaHallRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cms service

type CmsHandler interface {
	UserLogin(context.Context, *UserLoginReq, *UserLoginRsp) error
	UpdateMessage(context.Context, *UpdateMessageReq, *UpdateMessageRsp) error
	AllFilms(context.Context, *AllFilmsReq, *AllFilmsRsp) error
	UpdateFilm(context.Context, *UpdateFilmReq, *UpdateFilmRsp) error
	DeleteFilm(context.Context, *DeleteFilmReq, *DeleteFilmRsp) error
	AllUsers(context.Context, *AllUsersReq, *AllUsersRsp) error
	AllAdminUsers(context.Context, *AllAdminUsersReq, *AllAdminUsersRsp) error
	AllComments(context.Context, *AllCommentsReq, *AllCommentsRsp) error
	AllOrders(context.Context, *AllOrdersReq, *AllOrdersRsp) error
	AllAddress(context.Context, *AllAddressReq, *AllAddressRsp) error
	AddFilm(context.Context, *AddFilmReq, *AddFilmRsp) error
	AddAdminUser(context.Context, *AddAdminUserReq, *AddAdminUserRsp) error
	AddAddress(context.Context, *AddAddressReq, *AddAddressRsp) error
	UpdateAddress(context.Context, *UpdateAddressReq, *UpdateAddressRsp) error
	DeleteAddress(context.Context, *DeleteAddressReq, *DeleteAddressRsp) error
	DeleteAdminUser(context.Context, *DeleteAdminUserReq, *DeleteAdminUserRsp) error
	AllMovieHall(context.Context, *AllMovieHallReq, *AllMovieHallRsp) error
	AddMovieHall(context.Context, *AddMovieHallReq, *AddMovieHallRsp) error
	UpdateMovieHall(context.Context, *UpdateMovieHallReq, *UpdateMovieHallRsp) error
	DeleteMovieHall(context.Context, *DeleteMovieHallReq, *DeleteMovieHallRsp) error
	AllCinemaFilms(context.Context, *AllCinemaFilmsReq, *AllCinemaFilmsRsp) error
	AddCinemaFilm(context.Context, *AddCinemaFilmReq, *AddCinemaFilmRsp) error
	UpdateCinemaFilm(context.Context, *UpdateCinemaFilmReq, *UpdateCinemaFilmRsp) error
	DeleteCinemaFilm(context.Context, *DeleteCinemaFilmReq, *DeleteCinemaFilmRsp) error
	RegisterCinema(context.Context, *RegisterCinemaReq, *RegisterCinemaRsp) error
	AllCinemaHall(context.Context, *AllCinemaHallReq, *AllCinemaHallRsp) error
}

func RegisterCmsHandler(s server.Server, hdlr CmsHandler, opts ...server.HandlerOption) error {
	type cms interface {
		UserLogin(ctx context.Context, in *UserLoginReq, out *UserLoginRsp) error
		UpdateMessage(ctx context.Context, in *UpdateMessageReq, out *UpdateMessageRsp) error
		AllFilms(ctx context.Context, in *AllFilmsReq, out *AllFilmsRsp) error
		UpdateFilm(ctx context.Context, in *UpdateFilmReq, out *UpdateFilmRsp) error
		DeleteFilm(ctx context.Context, in *DeleteFilmReq, out *DeleteFilmRsp) error
		AllUsers(ctx context.Context, in *AllUsersReq, out *AllUsersRsp) error
		AllAdminUsers(ctx context.Context, in *AllAdminUsersReq, out *AllAdminUsersRsp) error
		AllComments(ctx context.Context, in *AllCommentsReq, out *AllCommentsRsp) error
		AllOrders(ctx context.Context, in *AllOrdersReq, out *AllOrdersRsp) error
		AllAddress(ctx context.Context, in *AllAddressReq, out *AllAddressRsp) error
		AddFilm(ctx context.Context, in *AddFilmReq, out *AddFilmRsp) error
		AddAdminUser(ctx context.Context, in *AddAdminUserReq, out *AddAdminUserRsp) error
		AddAddress(ctx context.Context, in *AddAddressReq, out *AddAddressRsp) error
		UpdateAddress(ctx context.Context, in *UpdateAddressReq, out *UpdateAddressRsp) error
		DeleteAddress(ctx context.Context, in *DeleteAddressReq, out *DeleteAddressRsp) error
		DeleteAdminUser(ctx context.Context, in *DeleteAdminUserReq, out *DeleteAdminUserRsp) error
		AllMovieHall(ctx context.Context, in *AllMovieHallReq, out *AllMovieHallRsp) error
		AddMovieHall(ctx context.Context, in *AddMovieHallReq, out *AddMovieHallRsp) error
		UpdateMovieHall(ctx context.Context, in *UpdateMovieHallReq, out *UpdateMovieHallRsp) error
		DeleteMovieHall(ctx context.Context, in *DeleteMovieHallReq, out *DeleteMovieHallRsp) error
		AllCinemaFilms(ctx context.Context, in *AllCinemaFilmsReq, out *AllCinemaFilmsRsp) error
		AddCinemaFilm(ctx context.Context, in *AddCinemaFilmReq, out *AddCinemaFilmRsp) error
		UpdateCinemaFilm(ctx context.Context, in *UpdateCinemaFilmReq, out *UpdateCinemaFilmRsp) error
		DeleteCinemaFilm(ctx context.Context, in *DeleteCinemaFilmReq, out *DeleteCinemaFilmRsp) error
		RegisterCinema(ctx context.Context, in *RegisterCinemaReq, out *RegisterCinemaRsp) error
		AllCinemaHall(ctx context.Context, in *AllCinemaHallReq, out *AllCinemaHallRsp) error
	}
	type Cms struct {
		cms
	}
	h := &cmsHandler{hdlr}
	return s.Handle(s.NewHandler(&Cms{h}, opts...))
}

type cmsHandler struct {
	CmsHandler
}

func (h *cmsHandler) UserLogin(ctx context.Context, in *UserLoginReq, out *UserLoginRsp) error {
	return h.CmsHandler.UserLogin(ctx, in, out)
}

func (h *cmsHandler) UpdateMessage(ctx context.Context, in *UpdateMessageReq, out *UpdateMessageRsp) error {
	return h.CmsHandler.UpdateMessage(ctx, in, out)
}

func (h *cmsHandler) AllFilms(ctx context.Context, in *AllFilmsReq, out *AllFilmsRsp) error {
	return h.CmsHandler.AllFilms(ctx, in, out)
}

func (h *cmsHandler) UpdateFilm(ctx context.Context, in *UpdateFilmReq, out *UpdateFilmRsp) error {
	return h.CmsHandler.UpdateFilm(ctx, in, out)
}

func (h *cmsHandler) DeleteFilm(ctx context.Context, in *DeleteFilmReq, out *DeleteFilmRsp) error {
	return h.CmsHandler.DeleteFilm(ctx, in, out)
}

func (h *cmsHandler) AllUsers(ctx context.Context, in *AllUsersReq, out *AllUsersRsp) error {
	return h.CmsHandler.AllUsers(ctx, in, out)
}

func (h *cmsHandler) AllAdminUsers(ctx context.Context, in *AllAdminUsersReq, out *AllAdminUsersRsp) error {
	return h.CmsHandler.AllAdminUsers(ctx, in, out)
}

func (h *cmsHandler) AllComments(ctx context.Context, in *AllCommentsReq, out *AllCommentsRsp) error {
	return h.CmsHandler.AllComments(ctx, in, out)
}

func (h *cmsHandler) AllOrders(ctx context.Context, in *AllOrdersReq, out *AllOrdersRsp) error {
	return h.CmsHandler.AllOrders(ctx, in, out)
}

func (h *cmsHandler) AllAddress(ctx context.Context, in *AllAddressReq, out *AllAddressRsp) error {
	return h.CmsHandler.AllAddress(ctx, in, out)
}

func (h *cmsHandler) AddFilm(ctx context.Context, in *AddFilmReq, out *AddFilmRsp) error {
	return h.CmsHandler.AddFilm(ctx, in, out)
}

func (h *cmsHandler) AddAdminUser(ctx context.Context, in *AddAdminUserReq, out *AddAdminUserRsp) error {
	return h.CmsHandler.AddAdminUser(ctx, in, out)
}

func (h *cmsHandler) AddAddress(ctx context.Context, in *AddAddressReq, out *AddAddressRsp) error {
	return h.CmsHandler.AddAddress(ctx, in, out)
}

func (h *cmsHandler) UpdateAddress(ctx context.Context, in *UpdateAddressReq, out *UpdateAddressRsp) error {
	return h.CmsHandler.UpdateAddress(ctx, in, out)
}

func (h *cmsHandler) DeleteAddress(ctx context.Context, in *DeleteAddressReq, out *DeleteAddressRsp) error {
	return h.CmsHandler.DeleteAddress(ctx, in, out)
}

func (h *cmsHandler) DeleteAdminUser(ctx context.Context, in *DeleteAdminUserReq, out *DeleteAdminUserRsp) error {
	return h.CmsHandler.DeleteAdminUser(ctx, in, out)
}

func (h *cmsHandler) AllMovieHall(ctx context.Context, in *AllMovieHallReq, out *AllMovieHallRsp) error {
	return h.CmsHandler.AllMovieHall(ctx, in, out)
}

func (h *cmsHandler) AddMovieHall(ctx context.Context, in *AddMovieHallReq, out *AddMovieHallRsp) error {
	return h.CmsHandler.AddMovieHall(ctx, in, out)
}

func (h *cmsHandler) UpdateMovieHall(ctx context.Context, in *UpdateMovieHallReq, out *UpdateMovieHallRsp) error {
	return h.CmsHandler.UpdateMovieHall(ctx, in, out)
}

func (h *cmsHandler) DeleteMovieHall(ctx context.Context, in *DeleteMovieHallReq, out *DeleteMovieHallRsp) error {
	return h.CmsHandler.DeleteMovieHall(ctx, in, out)
}

func (h *cmsHandler) AllCinemaFilms(ctx context.Context, in *AllCinemaFilmsReq, out *AllCinemaFilmsRsp) error {
	return h.CmsHandler.AllCinemaFilms(ctx, in, out)
}

func (h *cmsHandler) AddCinemaFilm(ctx context.Context, in *AddCinemaFilmReq, out *AddCinemaFilmRsp) error {
	return h.CmsHandler.AddCinemaFilm(ctx, in, out)
}

func (h *cmsHandler) UpdateCinemaFilm(ctx context.Context, in *UpdateCinemaFilmReq, out *UpdateCinemaFilmRsp) error {
	return h.CmsHandler.UpdateCinemaFilm(ctx, in, out)
}

func (h *cmsHandler) DeleteCinemaFilm(ctx context.Context, in *DeleteCinemaFilmReq, out *DeleteCinemaFilmRsp) error {
	return h.CmsHandler.DeleteCinemaFilm(ctx, in, out)
}

func (h *cmsHandler) RegisterCinema(ctx context.Context, in *RegisterCinemaReq, out *RegisterCinemaRsp) error {
	return h.CmsHandler.RegisterCinema(ctx, in, out)
}

func (h *cmsHandler) AllCinemaHall(ctx context.Context, in *AllCinemaHallReq, out *AllCinemaHallRsp) error {
	return h.CmsHandler.AllCinemaHall(ctx, in, out)
}
