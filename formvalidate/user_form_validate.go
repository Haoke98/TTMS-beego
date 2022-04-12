package formvalidate

import "github.com/gookit/validate"

// UserForm user 表单
type UserForm struct {
	Id             int    `form:"id"`
	Avatar         string `form:"avatar"`
	Username       string `form:"username" validate:"required"`
	FullName       string `form:"fullName" validate:"required"`
	FullNamePinyin string `form:"fullNamePinyin" validate:"required"`
	Mobile         string `form:"mobile" validate:"required"`
	Password       string `form:"password" validate:"required"`
	Status         int8   `form:"status"`
	Description    string `form:"description"`
	CreateTime     int    `form:"create_time"`
	UpdateTime     int    `form:"update_time"`
	DeleteTime     int    `form:"delete_time"`
	IsCreate       int    `form:"_create"`
}

// Messages 自定义验证返回消息
func (f UserForm) Messages() map[string]string {
	return validate.MS{
		"Title.required":          "用户名不能为空.",
		"Mobile.required":         "手机号不能为空.",
		"FullName.required":       "姓名不能为空.",
		"FullNamePinYin.required": "姓名（拼音）不能为空",
		"Summary.required":        "密码不能为空.",
	}
}
