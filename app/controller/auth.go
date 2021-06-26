package controller

import (
	"feedback/app/logic"
	"feedback/app/request"
	"feedback/app/util/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
	var req *request.UserRegisterRequest

	if err := r.Parse(&req); err != nil {
		response.Jump(r, 0, err.Error())
	}

	if _, err := logic.UserLogic.Register(req); err != nil {
		response.Jump(r, 0, err.Error())
	}

	response.Jump(r, 1, "注册成功", "/login")
}

func (*authController) Logout(r *ghttp.Request) {
	// ***
}
