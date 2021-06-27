package logic

import (
	"feedback/app/dao"
	"feedback/app/model"
	"feedback/app/request"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

type messageLogic struct{}

var MessageLogic = &messageLogic{}

// 提交留言
func (m *messageLogic) Message(r *ghttp.Request, req *request.MessagePostRequest) (*model.Message, error) {
	message := &model.Message{
		Uid:       0,
		Content:   req.Content,
		CreatedAt: gtime.Now(),
	}
	dao.Message.Insert(message)

	return message, nil
}

// 留言列表
func (m *messageLogic) Messages(r *ghttp.Request) (g.Map, error) {
	var messages []*model.Message
	per_size := 4

	total, _ := dao.Message.FindCount()

	page := r.GetPage(total, per_size)

	dao.Message.WithAll().Limit((page.CurrentPage-1)*per_size, per_size).OrderDesc("id").Structs(&messages)

	return g.Map{
		"total": total,
		"items": messages,
		"page":  page.GetContent(4),
	}, nil
}
