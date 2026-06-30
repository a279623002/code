package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	Id        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Username  string    `gorm:"column:username;size:64;not null;comment:用户名"`
	Phone     string    `gorm:"column:phone;size:20;uniqueIndex;comment:手机号"`
	Email     string    `gorm:"column:email;size:128;comment:邮箱"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}

// UserModel 用户数据访问层
type UserModel struct {
	db *gorm.DB
}

// NewUserModel 创建用户模型
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

// AutoMigrate 自动建表
func (m *UserModel) AutoMigrate() error {
	return m.db.AutoMigrate(&User{})
}

// Create 创建用户
func (m *UserModel) Create(ctx context.Context, user *User) error {
	return m.db.WithContext(ctx).Create(user).Error
}

// FindOne 根据 ID 查询用户
func (m *UserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	var user User
	if err := m.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
