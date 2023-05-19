package service

import (
	"context"

	v1 "module/api/module/v1"
	"module/internal/biz"
)

// ModuleService is a greeter service.
type ModuleService struct {
	v1.UnimplementedModuleServer

	uc *biz.ModuleUsecase
}

// NewModuleService new a greeter service.
func NewModuleService(uc *biz.ModuleUsecase) *ModuleService {
	return &ModuleService{uc: uc}
}

func (s *ModuleService) CreateModule(ctx context.Context, in *v1.CreateModuleRequest) (reply *v1.CreateModuleReply, err error) {

	return
}
func (s *ModuleService) UpdateModule(ctx context.Context, in *v1.UpdateModuleRequest) (reply *v1.UpdateModuleReply, err error) {

	return
}
func (s *ModuleService) DeleteModule(ctx context.Context, in *v1.DeleteModuleRequest) (reply *v1.DeleteModuleReply, err error) {

	return
}
func (s *ModuleService) GetModule(ctx context.Context, in *v1.GetModuleRequest) (reply *v1.GetModuleReply, err error) {
	reply = &v1.GetModuleReply{}
	return
}

func (s *ModuleService) ListModule(ctx context.Context, in *v1.PageInfoRequest) (reply *v1.ListModuleReply, err error) {

	return
}
