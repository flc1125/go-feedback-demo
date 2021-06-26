package logic

import (
	"errors"
	"feedback/app/dao"
	"feedback/app/model"
	"feedback/app/request"
	"feedback/app/service/password"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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

// 用户登录
func (u *userLogic) Login(r *ghttp.Request, req *request.UserLoginRequest) (*model.User, error) {
	// 查找用户数据
	user := &model.User{}
	err := dao.User.Struct(user, g.Map{
		"username": req.Username,
		"status":   1,
	})
	if err != nil {
		return nil, errors.New("用户名错误")
	}

	// 密码验证
	if !password.Verify(req.Password, user.Password) {
		return nil, errors.New("用户密码错误")
	}

	// 登录状态：Session
	r.Session.Set("user", user)

	return user, nil
}
