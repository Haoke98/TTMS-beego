package controllers

import (
	"TTMS/formvalidate"
	"TTMS/global"
	"TTMS/global/response"
	"TTMS/models"
	"TTMS/services"
	"TTMS/utils"
	"encoding/base64"
	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
	"strconv"
	"strings"
)

// TrainPlanController  培训计划控制器
type TrainPlanController struct {
	baseController
}

// Index 培训计划管理-首页
func (auc *TrainPlanController) Index() {
	var trainService services.TrainPlanService
	data, pagination := trainService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	auc.Data["data"] = data
	auc.Data["paginate"] = pagination

	auc.Layout = "public/base.html"
	auc.TplName = "train_plan/index.html"
}

type PlanDTO1 struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Cover      string `json:"cover"`
	Status     int    `json:"status"`
	Favor      bool   `json:"favor"`
	Applicants int    `json:"applicants"`
	Quota      int    `json:"quota"`
	Start      string `json:"start"`
	End        string `json:"end"`
}

func (c *TrainPlanController) Get() {
	var trainPlanService services.TrainPlanService
	var quotaService services.QuotaService
	user, ok := c.Ctx.Input.Session(global.LOGIN_USER).(models.User)
	if ok {
		data, pagination := trainPlanService.GetPaginateData(admin["per_page"].(int), gQueryParams)
		var plans []PlanDTO1
		for _, plan := range data {
			plans = append(plans, PlanDTO1{
				Id:         plan.Id,
				Title:      plan.Title,
				Summary:    "想不想通过专业的培训达到自己的教育梦，提高自己的教师技能，课堂效率？那还愣着干嘛？赶快行动啊！",
				Cover:      "",
				Status:     0,
				Favor:      trainPlanService.IsFavor(plan.Id, user.Id),
				Applicants: 92,
				Quota:      quotaService.GetTotalCount(plan.Id),
				Start:      plan.RegistrationStartedAt.Format("06/01/02 15:04"),
				End:        plan.RegistrationStartedAt.Format("06/01/02 15:04")})
		}
		c.Data["json"] = map[string]interface{}{
			"plans":      plans,
			"pagination": pagination,
		}
		c.ServeJSON()
	}
}

// Add 培训计划管理-添加界面
func (auc *TrainPlanController) Add() {
	var (
		adminRoleService services.AdminRoleService
		trainPlanService services.TrainPlanService
	)

	roles := adminRoleService.GetAllData()
	tempPlan := trainPlanService.GetTempTrainPlan()
	auc.Data["roles"] = roles
	auc.Data["data"] = tempPlan
	auc.Layout = "public/base.html"
	auc.TplName = "train_plan/edit.html"
}

// Create 用户管理-添加界面
func (auc *TrainPlanController) Create() {
	var trainForm formvalidate.TrainPlanForm
	if err := auc.ParseForm(&trainForm); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}
	v := validate.Struct(trainForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), auc.Ctx)
	}

	//账号验重
	var trainService services.TrainPlanService
	if trainService.IsExistName(strings.TrimSpace(trainForm.Title), 0) {
		response.ErrorWithMessage("同名培训计划已存在！", auc.Ctx)
	}

	insertID, _ := trainService.Create(&trainForm)

	url := global.URL_BACK
	if trainForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Edit 系统管理-用户管理-修改界面
func (auc *TrainPlanController) Edit() {
	id, _ := auc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", auc.Ctx)
	}

	var (
		trainService services.TrainPlanService
	)

	train := trainService.GetById(id)
	if train == nil {
		response.ErrorWithMessage("Not Found Info By Id.", auc.Ctx)
	}
	auc.Data["data"] = train

	auc.Layout = "public/base.html"
	auc.TplName = "train_plan/edit.html"
}

// Update 系统管理-用户管理-修改
func (auc *TrainPlanController) Update() {
	var trainForm formvalidate.TrainPlanForm
	if err := auc.ParseForm(&trainForm); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}

	if trainForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", auc.Ctx)
	}

	v := validate.Struct(trainForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), auc.Ctx)
	}

	//账号验重
	var trainService services.TrainPlanService
	if trainService.IsExistName(strings.TrimSpace(trainForm.Title), trainForm.Id) {
		response.ErrorWithMessage("账号已经存在", auc.Ctx)
	}

	num := trainService.Update(&trainForm)

	if num > 0 {
		response.Success(auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Enable 启用
func (auc *TrainPlanController) Enable() {
	idStr := auc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		auc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择启用的用户.", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Disable 禁用
func (auc *TrainPlanController) Disable() {
	idStr := auc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		auc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择禁用的用户.", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Del 删除
func (auc *TrainPlanController) Del() {
	idStr := auc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		auc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", auc.Ctx)
	}

	noDeletionID := new(models.AdminUser).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"的数据无法删除!", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	count := adminUserService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Profile 系统管理-个人资料
func (auc *TrainPlanController) Profile() {
	auc.Layout = "public/base.html"
	auc.TplName = "admin_user/profile.html"
}

// UpdateNickName 系统管理-个人资料-修改昵称
func (auc *TrainPlanController) UpdateNickName() {
	id, err := auc.GetInt("id")
	nickname := strings.TrimSpace(auc.GetString("nickname"))

	if nickname == "" || err != nil {
		response.ErrorWithMessage("参数错误", auc.Ctx)
	}

	// 验证是否是登陆用户，这里也可不用提供的id，使用登陆的id即可
	if loginUser.Id != id {
		response.ErrorWithMessage("数据非法", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.UpdateNickName(id, nickname)

	if num > 0 {
		//修改成功后，更新session的登录用户信息
		loginAdminUser := adminUserService.GetAdminUserById(id)
		auc.SetSession(global.LOGIN_ADMIN_USER, *loginAdminUser)
		response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// UpdatePassword 系统管理-个人资料-修改密码
func (auc *TrainPlanController) UpdatePassword() {
	id, err := auc.GetInt("id")
	password := auc.GetString("password")
	newPassword := auc.GetString("new_password")
	reNewPassword := auc.GetString("renew_password")

	// 验证是否是登陆用户，这里也可不用提供的id，使用登陆的id即可
	if loginUser.Id != id {
		response.ErrorWithMessage("数据非法", auc.Ctx)
	}

	if err != nil || password == "" || newPassword == "" || reNewPassword == "" {
		response.ErrorWithMessage("Bad Parameter.", auc.Ctx)
	}

	if newPassword != reNewPassword {
		response.ErrorWithMessage("两次输入的密码不一致.", auc.Ctx)
	}

	if password == newPassword {
		response.ErrorWithMessage("新密码与旧密码一致，无需修改", auc.Ctx)
	}

	loginUserPassword, err := base64.StdEncoding.DecodeString(loginUser.Password)

	if err != nil {
		response.ErrorWithMessage("err:"+err.Error(), auc.Ctx)
	}

	if !utils.PasswordVerify(password, string(loginUserPassword)) {
		response.ErrorWithMessage("当前密码不正确", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.UpdatePassword(id, newPassword)
	if num > 0 {
		response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// UpdateAvatar 系统管理-个人资料-修改头像
func (auc *TrainPlanController) UpdateAvatar() {
	_, _, err := auc.GetFile("avatar")
	if err != nil {
		response.ErrorWithMessage("上传头像错误"+err.Error(), auc.Ctx)
	}

	var (
		attachmentService services.AttachmentService
		adminUserService  services.AdminUserService
	)
	attachmentInfo, err := attachmentService.Upload(auc.Ctx, "avatar", loginUser.Id, 0)
	if err != nil || attachmentInfo == nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	} else {
		//头像上传成功，更新用户的avatar头像信息
		num := adminUserService.UpdateAvatar(loginUser.Id, attachmentInfo.Url)
		if num > 0 {
			//修改成功后，更新session的登录用户信息
			loginAdminUser := adminUserService.GetAdminUserById(loginUser.Id)
			auc.SetSession(global.LOGIN_ADMIN_USER, *loginAdminUser)
			response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, auc.Ctx)
		} else {
			response.Error(auc.Ctx)
		}
	}

}
