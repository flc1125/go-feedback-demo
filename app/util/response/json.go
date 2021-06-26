package response

import "github.com/gogf/gf/net/ghttp"

type JsonResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func Json(r *ghttp.Request, status int, msg string, data ...interface{}) {
	responseData := interface{}(nil)

	if len(data) > 0 {
		responseData = data[0]
	}

	r.Response.WriteJson(JsonResponse{
		Status: status,
		Msg:    msg,
		Data:   responseData,
	})
}
