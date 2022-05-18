package formvalidate

import (
	"github.com/gookit/validate"
	"time"
)

// TrainPlanForm 培训计划表单
type TrainPlanForm struct {
	Id                    int       `form:"id"`
	Title                 string    `form:"title" validate:"required"`
	Summary               string    `form:"summary" validate:"required"`
	RegistrationStartedAt time.Time `form:"registrationStartedAt" validate:"required"`
	RegistrationEndAt     time.Time `form:"registrationEndAt" validate:"required"`
	PersonInCharge        string    `form:"personInCharge" validate:"required"`
	Status                int       `form:"status"`
	IsCreate              int       `form:"_create"`
}

// Messages 自定义验证返回消息
func (f TrainPlanForm) Messages() map[string]string {
	return validate.MS{
		"Name.required":                  "请填写标题.",
		"Code.required":                  "请填简介.",
		"RegistrationStartedAt.required": "请选择报名开始时间.",
		"RegistrationEndAt.required":     "请选择报名结束时间.",
		"PersonInCharge.required":        "请填写负责人姓名.",
	}
}
