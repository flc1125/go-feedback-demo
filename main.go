package main

import (
	_ "feedback/boot"
	_ "feedback/router"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

func main() {
	gtime.SetTimeZone("PRC")
	g.Server().Run()
}
