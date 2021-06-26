package response

import "github.com/gogf/gf/frame/g"
import "github.com/gogf/gf/net/ghttp"

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

	// URL
	if url, ok := options[0]; ok {
		j.Url = url
	} else {
		if status == 1 {
			j.Url = r.GetReferer() // 返回来源页
		} else {
			j.Url = "javascript:history.back(-1);" // js返回上一页
		}
	}

	// Wait
	if wait, ok := options[1]; ok {
		j.Wait = wait
	} else {
		if status == 1 {
			j.Wait = 1
		} else {
			j.Wait = 3
		}
	}

	r.Response.WriteTpl("public/jump.html", g.Map{
		"j": j,
	})

	r.Exit()
}
