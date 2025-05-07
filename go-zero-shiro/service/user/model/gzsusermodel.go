package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GzsUserModel = (*customGzsUserModel)(nil)

type (
	// GzsUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGzsUserModel.
	GzsUserModel interface {
		gzsUserModel
		FindOneByNumber(ctx context.Context, name string) (*GzsUser, error)
	}

	customGzsUserModel struct {
		*defaultGzsUserModel
	}
)

// NewGzsUserModel returns a model for the database table.
func NewGzsUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GzsUserModel {
	return &customGzsUserModel{
		defaultGzsUserModel: newGzsUserModel(conn, c, opts...),
	}
}

func (m *defaultGzsUserModel) FindOneByNumber(ctx context.Context, name string) (*GzsUser, error) {
	goZeroShiroGzsUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroShiroGzsUserIdPrefix, name)
	var resp GzsUser
	err := m.QueryRowCtx(ctx, &resp, goZeroShiroGzsUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", gzsUserRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, name)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
