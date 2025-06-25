package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, g *gorm.DB) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, g),
	}
}

