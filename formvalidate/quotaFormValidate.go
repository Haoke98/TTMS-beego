package formvalidate

import (
	"github.com/gookit/validate"
)

// QuotaForm 培训计划表单
type QuotaForm struct {
	PlanId       int `form:"planId" validate:"required"`
	UniversityId int `form:"universityId" validate:"required"`
	Quota        int `form:"quota" validate:"required"`
}

// Messages 自定义验证返回消息
func (f QuotaForm) Messages() map[string]string {
	return validate.MS{
		"PlanId.required":     "请提供培训计划ID.",
		"University.required": "请提供高校ID.",
		"Quota.required":      "请提供报名名额数量.",
	}
}
