package request

// 提交注册参数
type UserRegisterRequest struct {
	Username        string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:min到:max位"`
	Password        string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	PasswordConfirm string `p:"password_confirm" v:"required|length:6,30|same:password#请确认密码|密码长度不够|两次密码不一致"`
}

// 提交登录参数
type UserLoginRequest struct {
	Username string `p:"username"  v:"required#请输入账号"`
	Password string `p:"password" v:"required#请输入密码"`
}

type UserStatusRequest struct {
	Id string `p:"id"  v:"required#ID错误"`
	S  int    `p:"s" v:"required#参数错误"`
}

type UserBlackRequest struct {
	Id string `p:"id"  v:"required#ID错误"`
	S  int    `p:"s" v:"required#参数错误"`
}
