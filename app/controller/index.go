package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var IndexController = indexController{}

type indexController struct{}

func (*indexController) Index(r *ghttp.Request) {
	// (&user.Entity{
	// 	Username:  "flc1125",
	// 	Password:  "123456",
	// 	CreatedAt: gtime.Now(),
	// 	UpdatedAt: gtime.Now(),
	// 	IsAdmin:   1,
	// }).Insert()

	r.Response.WriteTpl("index.html", g.Map{})
}
