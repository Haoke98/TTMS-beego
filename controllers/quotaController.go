package controllers

import (
	"beego-admin/global/response"
	"beego-admin/services"
)

type QuotaController struct {
	baseController
}

// Index 用户管理-首页
func (c *QuotaController) Index() {
	id, _ := c.GetInt("id", -1)
	planId, _ := c.GetInt("plan_id", -1)
	if id <= 0 && planId <= 0 {
		response.ErrorWithMessage("Param is error.", c.Ctx)
	}
	var service services.QuotaService
	data, pagination := service.GetPaginateData(admin["per_page"].(int), gQueryParams)
	c.Data["data"] = data
	c.Data["paginate"] = pagination
	c.Layout = "public/base_modal.html"
	c.TplName = "quota/index.html"
}
