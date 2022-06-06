package controllers

import (
	"TTMS/formvalidate"
	"TTMS/global/response"
	"TTMS/services"
	"github.com/gookit/validate"
)

type PetitionController struct {
	baseController
}

// Index 用户管理-首页
func (c *PetitionController) Index() {
	id, _ := c.GetInt("id", -1)
	planId, _ := c.GetInt("plan_id", -1)
	if id <= 0 && planId <= 0 {
		response.ErrorWithMessage("Param is error.", c.Ctx)
	} else {
		if planId == -1 && id > 0 {
			planId = id
		}
	}
	var (
		petitionService  services.PetitionService
		trainPlanService services.TrainPlanService
	)
	trainPlan := trainPlanService.GetById(planId)
	petitions, pagination := petitionService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	c.Data["petitions"] = petitions
	c.Data["plan"] = trainPlan
	c.Data["paginate"] = pagination
	c.Layout = "public/base_modal.html"
	c.TplName = "petition/index.html"
}

// Update 系统管理-用户管理-修改
func (c *PetitionController) Update() {
	var form formvalidate.QuotaForm
	if err := c.ParseForm(&form); err != nil {
		response.ErrorWithMessage(err.Error(), c.Ctx)
	}
	v := validate.Struct(form)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), c.Ctx)
	}

	//账号验重
	var service services.QuotaService
	num := service.CreateOrUpdate(&form)
	//if service.IsExistName(strings.TrimSpace(form.Name), form.Id) {
	//	response.ErrorWithMessage("账号已经存在", c.Ctx)
	//}
	//
	//num := service.Update(&form)

	if num > 0 {
		response.Success(c.Ctx)
	} else {
		response.Error(c.Ctx)
	}
}
