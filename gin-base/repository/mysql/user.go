package mysql

import (
	"crypto/md5"
	"edu-admin-api/models"

	"encoding/hex"
	"errors"
)

const secret = "abc"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

// 检查用户是否存在
func CheckUserExist(username string) (err error) {
	var u models.User
	var count int64

	// 迁移 schema（自动创建'user'表）
	db.AutoMigrate(&u)

	//请求返回记录数
	if err := db.Model(&u).Where("username = ?", username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// 插入数据
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	db.Create(&user)
	return
}

// 登录
func Login(user *models.User) (u *models.User, err error) {
	oPassword := user.Password
	db.Select([]string{"user.id", "user.user_id", "user.username", "user.password"}).Where("user.username = ?", user.Username).Find(&user)
	if user.ID == 0 {
		return nil, ErrorUserNotExist
	}
	if db.Error != nil {
		//查询数据库失败
		return nil, err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return nil, ErrorInvalidPassword
	}
	return user, err
}

// MD5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
