package middleware

import (
	"feedback/app/dao"
	"feedback/app/model"
	"feedback/app/util/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type auth struct{}

var Auth = auth{}

func (*auth) Init(r *ghttp.Request) {
	v := r.Session.Get("user")

	if v == nil {
		r.Session.Remove("user")
		r.Middleware.Next()
		return
	}
	vv, _ := v.(*model.User)

	// 登录信息
	user := &model.User{}
	err := dao.User.Struct(user, g.Map{
		"id":     vv.Id,
		"status": 1,
	})

	if err != nil {
		r.Session.Remove("user")
		r.Middleware.Next()
		return
	}

	// 更新 Session
	r.Session.Set("user", user)

	r.Middleware.Next()
}

func (*auth) Check(r *ghttp.Request) {
	v := r.Session.Get("user")

	if v == nil {
		r.Session.Remove("user")
		response.Jump(r, 0, "请先登录", "/login")
	}

	r.Middleware.Next()
}
