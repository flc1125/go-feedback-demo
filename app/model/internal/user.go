// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `orm:"id,primary"      json:"id"`        //
	Username  string      `orm:"username,unique" json:"username"`  // 用户名
	Password  string      `orm:"password"        json:"password"`  // 密码
	CreatedAt *gtime.Time `orm:"created_at"      json:"createdAt"` // 注册时间
	UpdatedAt *gtime.Time `orm:"updated_at"      json:"updatedAt"` // 更新时间
	IsAdmin   int         `orm:"is_admin"        json:"isAdmin"`   // 是否为管理员
	Status    int         `orm:"status"          json:"status"`    // 是否可用
	Black     int         `orm:"black"           json:"black"`     // 是否为黑名单
}
