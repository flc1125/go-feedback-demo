package router

import (
	"feedback/app/api"
	"feedback/app/controller"
	"feedback/app/controller/u"
	"feedback/app/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth.Init)

		group.GET("/", controller.IndexController.Index)

		// 授权相关
		group.GET("/login", controller.AuthController.Login)
		group.POST("/login", controller.AuthController.DoLogin)
		group.GET("/register", controller.AuthController.Register)
		group.POST("/register", controller.AuthController.DoRegister)
		group.GET("/logout", controller.AuthController.Logout)

		// 用户中心
		group.Group("/u", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth.Check)

			group.GET("/", u.UUserController.Index)
			group.GET("/users", u.UUserController.Users)
			group.GET("/user-status", u.UUserController.UserStatus)
			group.GET("/user-black", u.UUserController.UserBlack)
		})

		// 用户留言
		group.POST("/message", controller.MessageController.Message).Middleware(middleware.Auth.Check)

		group.ALL("/hello", api.Hello)
	})
}
