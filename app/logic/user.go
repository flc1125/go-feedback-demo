package logic

import (
	"errors"
	"feedback/app/dao"
	"feedback/app/model"
	"feedback/app/request"
	"feedback/app/service/password"

	"github.com/gogf/gf/os/gtime"
)

type userLogic struct{}

var UserLogic = &userLogic{}

// 注册
func (u *userLogic) Register(r *request.UserRegisterRequest) (*model.User, error) {
	// 验证是否可注册
	if !u.checkUsername(r.Username) {
		return nil, errors.New("用户名已存在")
	}

	// 入库
	user := &model.User{
		Username:  r.Username,
		Password:  password.MustHash(r.Password),
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
		IsAdmin:   0,
		Status:    1,
	}
	dao.User.Insert(user)

	return user, nil
}

// 验证是否有重复的用户名
func (u *userLogic) checkUsername(username string) bool {
	if c, err := dao.User.FindCount("username", username); err != nil || c >= 1 {
		return false
	}

	return true
}
