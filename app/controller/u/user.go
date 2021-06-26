package u

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var UUserController = uUserController{}

type uUserController struct{}

func (*uUserController) Users(r *ghttp.Request) {
	r.Response.WriteTpl("u/users.html", g.Map{})
}
