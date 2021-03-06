package logic

import (
	"errors"
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
	var user *model.User
	r.Session.GetVar("user").Struct(&user)

	// 黑名单不可留言
	if user.Black == 1 {
		return nil, errors.New("黑名单用户不可留言")
	}

	message := &model.Message{
		Uid:       user.Id,
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
