package formvalidate

import (
	"github.com/gookit/validate"
)

// UniversityForm 培训计划表单
type UniversityForm struct {
	Id       int    `form:"id"`
	Name     string `form:"name" validate:"required"`
	Code     string `form:"code" validate:"required"`
	Badge    string `form:"badge"`
	IsCreate int    `form:"_create"`
}

// Messages 自定义验证返回消息
func (f UniversityForm) Messages() map[string]string {
	return validate.MS{
		"Name.required": "请填写院校名称.",
		"Code.required": "请填写院校编号.",
	}
}
