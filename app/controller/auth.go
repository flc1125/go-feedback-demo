package controller

import (
	"feedback/app/model/user"
	"feedback/app/service/password"
	"feedback/app/util/response"
	"feedback/app/validate"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

var AuthController = authController{}

type authController struct{}

func (*authController) Login(r *ghttp.Request) {
	r.Response.WriteTpl("login.html", g.Map{})
}

func (*authController) Register(r *ghttp.Request) {
	r.Response.WriteTpl("register.html", g.Map{})
}

func (*authController) DoRegister(r *ghttp.Request) {
	var req *validate.UserRegisterReq

	if err := r.Parse(&req); err != nil {
		response.Jump(r, 0, err.Error())
	}

	(&user.Entity{
		Username:  req.Username,
		Password:  password.MustHash(req.Password),
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
		IsAdmin:   0,
	}).Insert()

	jump.Success(r, "/login", "注册成功")
}

func (*authController) Logout(r *ghttp.Request) {
	// ***
}
