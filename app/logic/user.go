package logic

import "feedback/app/model/user"

type User struct{}

var userLogic = &User{}

func (user *User) Register(r *ghttp.Request) (*user, error) {

	(&user.Entity{
		Username:  req.Username,
		Password:  password.MustHash(req.Password),
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
		IsAdmin:   0,
	}).Insert()

	return u
}
