package u

import (
	"feedback/app/logic"
	"feedback/app/request"
	"feedback/app/util/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var UUserController = uUserController{}

type uUserController struct{}

func (*uUserController) Index(r *ghttp.Request) {
	r.Response.WriteTpl("u/index.html", g.Map{})
}

func (*uUserController) Users(r *ghttp.Request) {
	users, err := logic.UserLogic.Users(r)
	if err != nil {
		response.Jump(r, 0, err.Error())
	}

	// r.Response.WriteJson(users)

	r.Response.WriteTpl("u/users.html", g.Map{
		"users": users,
	})
}

func (*uUserController) UserStatus(r *ghttp.Request) {
	var req *request.UserStatusRequest

	if err := r.Parse(&req); err != nil {
		response.Jump(r, 0, err.Error())
	}

	logic.UserLogic.UserStatus(req)

	response.Jump(r, 1, "操作成功")
}

func (*uUserController) UserBlack(r *ghttp.Request) {
	var req *request.UserBlackRequest

	if err := r.Parse(&req); err != nil {
		response.Jump(r, 0, err.Error())
	}

	logic.UserLogic.UserBlack(req)

	response.Jump(r, 1, "操作成功")
}
