package request

type MessagePostRequest struct {
	Content string `p:"content"  v:"required|length:4,100#请说点什么……|内容长度为:min到:max位"`
}
