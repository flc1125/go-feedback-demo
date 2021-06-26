package response

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type JumpMap struct {
	Status  int
	Message string
	Wait    int
	Url     string
}

func Jump(r *ghttp.Request, status int, message string, options ...interface{}) {
	j := &JumpMap{
		Status:  status,
		Message: message,
	}

	// 默认值
	// URL
	if status == 1 {
		j.Url = r.GetReferer() // 返回来源页
	} else {
		j.Url = "javascript:history.back(-1);" // js返回上一页
	}
	// Wait
	if status == 1 {
		j.Wait = 1
	} else {
		j.Wait = 3
	}

	// 定制化 URL
	if len(options) >= 1 {
		j.Url = options[0].(string)
	}

	// 定制化 Wait
	if len(options) >= 2 {
		j.Wait = options[1].(int)
	}

	r.Response.WriteTpl("public/jump.html", g.Map{
		"j": j,
	})

	r.Exit()
}
