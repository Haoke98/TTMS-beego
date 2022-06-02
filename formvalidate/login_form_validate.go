package formvalidate

import "github.com/gookit/validate"

// LoginForm login 表单
type LoginForm struct {
	Username  string `form:"username" json:"username" validate:"required"`
	Password  string `form:"password" json:"password" validate:"required"`
	Captcha   string `form:"captcha" json:"captcha" `
	CaptchaId string `form:"captchaId" json:"captchaId" `
	Remember  bool   `form:"remember" json:"remember" validate:"required"`
}

// Messages 自定义验证返回消息
func (f LoginForm) Messages() map[string]string {
	return validate.MS{
		"Username.required":  "用户名不能为空.",
		"Password.required":  "密码不能为空.",
		"Captcha.required":   "验证码内容不能为空.",
		"CaptchaId.required": "验证码ID不能为空.",
		"remember.required":  "是否记住登陆状态不能为空.",
	}
}
