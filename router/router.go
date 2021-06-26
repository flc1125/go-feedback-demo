package router

import (
	"feedback/app/api"
	"feedback/app/controller"
	"feedback/app/controller/u"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", controller.IndexController.Index)

		// 授权相关
		group.GET("/login", controller.AuthController.Login)
		group.GET("/register", controller.AuthController.Register)
		group.POST("/register", controller.AuthController.DoRegister)

		// 用户中心
		group.Group("/u", func(group *ghttp.RouterGroup) {
			group.GET("/users", u.UUserController.Users)
		})

		group.ALL("/hello", api.Hello)
	})
}
