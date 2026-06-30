package gorm

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config 数据库配置
type Config struct {
	DataSource      string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

// New 创建 gorm 连接
func New(cfg Config) (*gorm.DB, error) {
	if cfg.MaxIdleConns <= 0 {
		cfg.MaxIdleConns = 10
	}
	if cfg.MaxOpenConns <= 0 {
		cfg.MaxOpenConns = 100
	}
	if cfg.ConnMaxLifetime <= 0 {
		cfg.ConnMaxLifetime = 3600
	}

	db, err := gorm.Open(mysql.Open(cfg.DataSource), &gorm.Config{
		Logger: NewLogger(),
	})
	if err != nil {
		return nil, fmt.Errorf("gorm open failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get sql db failed: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	return db, nil
}

// zeroLogger gorm logger 适配 go-zero logx
type zeroLogger struct {
	level logger.LogLevel
}

// NewLogger 创建 gorm logger
func NewLogger() logger.Interface {
	return &zeroLogger{level: logger.Info}
}

func (l *zeroLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

func (l *zeroLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Infof(msg, data...)
}

func (l *zeroLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Infof("[WARN] "+msg, data...)
}

func (l *zeroLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Errorf(msg, data...)
}

func (l *zeroLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= 0 {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		logx.WithContext(ctx).Errorf("[%.3fms] [rows:%v] %s error: %v", float64(elapsed.Nanoseconds())/1e6, rows, sql, err)
		return
	}
	if l.level >= logger.Info {
		logx.WithContext(ctx).Infof("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}
