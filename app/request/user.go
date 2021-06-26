package request

type UserRegisterRequest struct {
	Username        string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:min到:max位"`
	Password        string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	PasswordConfirm string `p:"password_confirm" v:"required|length:6,30|same:password#请确认密码|密码长度不够|两次密码不一致"`
}
