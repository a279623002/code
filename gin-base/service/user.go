package service

import (
	"edu-admin-api/models"
	"edu-admin-api/pkg/snowflake"
	"edu-admin-api/repository/mysql"
	"fmt"
)

// 用户注册
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2.生成UID
	userID, _ := snowflake.GetID()

	// 3.构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存数据库
	return mysql.InsertUser(user)
}

// 用户登录
func Login(p map[string]interface{}) (u *models.User, err error) {
	user := &models.User{
		Username: fmt.Sprintf("%s", p["username"]),
		Password: fmt.Sprintf("%s", p["password"]),
	}
	u, err = mysql.Login(user)
	if err != nil {
		return nil, err
	}
	return u, err
}
