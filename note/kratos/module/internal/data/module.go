package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"module/internal/biz"
)

type Module struct {
	Id   uint64 `gorm:"primarykey;type:int;column:id;not null;"`
	CId  uint64 `gorm:"type:int;column:c_id;not null;"`
	Name string `gorm:"type:string;column:id;not null;"`
}

func (Module) TableName() string {
	return "user"
}

type moduleRepo struct {
	data *Data
	log  *log.Helper
}

func (m moduleRepo) CreateModule(ctx context.Context) error {
	return nil
}

func (m moduleRepo) UpdateModule(ctx context.Context) error {
	return nil
}

func (m moduleRepo) DeleteModule(ctx context.Context) error {
	return nil
}

func (m moduleRepo) GetModule(ctx context.Context) error {
	return nil
}

func (m moduleRepo) ListModule(ctx context.Context) error {
	return nil
}

// NewModuleRepo .
func NewModuleRepo(data *Data, logger log.Logger) biz.ModuleRepo {
	return &moduleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
