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
		group.POST("/login", controller.AuthController.DoLogin)
		group.GET("/register", controller.AuthController.Register)
		group.POST("/register", controller.AuthController.DoRegister)

		// 用户中心
		group.Group("/u", func(group *ghttp.RouterGroup) {
			group.GET("/users", u.UUserController.Users)
		})

		// 用户留言
		group.POST("/message", controller.MessageController.Message)

		group.ALL("/hello", api.Hello)
	})
}
