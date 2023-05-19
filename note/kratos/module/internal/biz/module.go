package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ModuleRepo interface {
	CreateModule(ctx context.Context) error
	UpdateModule(ctx context.Context) error
	DeleteModule(ctx context.Context) error
	GetModule(ctx context.Context) error
	ListModule(ctx context.Context) error
}

type ModuleUsecase struct {
	repo ModuleRepo
	log  *log.Helper
}

// NewModuleUsecase
func NewModuleUsecase(repo ModuleRepo, logger log.Logger) *ModuleUsecase {
	return &ModuleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (m *ModuleUsecase) CreateModuel(ctx context.Context) (err error) {
	return m.CreateModuel(ctx)
}
func (m *ModuleUsecase) UpdateModule(ctx context.Context) (err error) {
	return m.UpdateModule(ctx)
}
func (m *ModuleUsecase) DeleteModule(ctx context.Context) (err error) {
	return m.DeleteModule(ctx)
}
func (m *ModuleUsecase) GetModule(ctx context.Context) (err error) {
	return m.GetModule(ctx)
}
func (m *ModuleUsecase) ListModule(ctx context.Context) (err error) {
	return m.ListModule(ctx)
}
