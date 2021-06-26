package main

import (
	_ "feedback/boot"
	_ "feedback/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
