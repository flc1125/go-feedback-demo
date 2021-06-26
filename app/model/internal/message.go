// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Message is the golang structure for table message.
type Message struct {
	Id        int         `orm:"id,primary" json:"id"`        //
	Uid       int         `orm:"uid"        json:"uid"`       // 用户ID
	Message   string      `orm:"message"    json:"message"`   // 留言内容
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
}