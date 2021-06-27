package controller

import (
	"feedback/app/logic"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var IndexController = indexController{}

type indexController struct{}

func (*indexController) Index(r *ghttp.Request) {

	messages, _ := logic.MessageLogic.Messages(r)

	r.Response.WriteTpl("index.html", g.Map{
		"messages": messages,
	})
}
