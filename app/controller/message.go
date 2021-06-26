package controller

import (
	"feedback/app/logic"
	"feedback/app/request"
	"feedback/app/util/response"

	"github.com/gogf/gf/net/ghttp"
)

type messageController struct{}

var MessageController = &messageController{}

func (m *messageController) Message(r *ghttp.Request) {
	var req *request.MessagePostRequest

	if err := r.Parse(&req); err != nil {
		response.Jump(r, 0, err.Error())
	}

	if _, err := logic.MessageLogic.Message(r, req); err != nil {
		response.Jump(r, 0, err.Error())
	}

	response.Jump(r, 1, "留言成功")
}
