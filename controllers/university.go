package controllers

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/models"
	"beego-admin/services"
	"beego-admin/utils"
	"encoding/base64"
	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
	"strconv"
	"strings"
)

// UniversityController  高级院校控制器
type UniversityController struct {
	baseController
}

// Index 用户管理-首页
func (auc *UniversityController) Index() {
	var service services.UniversityService
	data, pagination := service.GetPaginateData(admin["per_page"].(int), gQueryParams)
	auc.Data["data"] = data
	auc.Data["paginate"] = pagination

	auc.Layout = "public/base.html"
	auc.TplName = "university/index.html"
}

// Add 用户管理-添加界面
func (auc *UniversityController) Add() {
	var adminRoleService services.AdminRoleService
	roles := adminRoleService.GetAllData()

	auc.Data["roles"] = roles
	auc.Layout = "public/base.html"
	auc.TplName = "university/add.html"
}

// Create 用户管理-添加界面
func (auc *UniversityController) Create() {
	var form formvalidate.UniversityForm
	if err := auc.ParseForm(&form); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}
	v := validate.Struct(form)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), auc.Ctx)
	}

	//账号验重
	var service services.UniversityService
	if service.IsExistName(strings.TrimSpace(form.Name), 0) {
		response.ErrorWithMessage("同名高级院校已存在！", auc.Ctx)
	}
	if service.IsExistCode(strings.TrimSpace(form.Code), 0) {
		response.ErrorWithMessage("同编号高级院校已存在！", auc.Ctx)
	}
	//处理图片上传
	_, _, err := auc.GetFile("badge")
	if err == nil {
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(auc.Ctx, "badge", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), auc.Ctx)
		} else {
			form.Badge = attachmentInfo.Url
		}
	}
	insertID := service.Create(&form)

	url := global.URL_BACK
	if form.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Edit 系统管理-用户管理-修改界面
func (auc *UniversityController) Edit() {
	id, _ := auc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", auc.Ctx)
	}

	var (
		service services.UniversityService
	)

	obj := service.GetUniversityById(id)
	if obj == nil {
		response.ErrorWithMessage("Not Found Info By Id.", auc.Ctx)
	}
	auc.Data["data"] = obj

	auc.Layout = "public/base.html"
	auc.TplName = "university/edit.html"
}

// Update 高校管理-修改-更新
func (auc *UniversityController) Update() {
	var form formvalidate.UniversityForm
	if err := auc.ParseForm(&form); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}

	if form.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", auc.Ctx)
	}

	v := validate.Struct(form)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), auc.Ctx)
	}

	//账号验重
	var service services.UniversityService
	if service.IsExistName(strings.TrimSpace(form.Name), form.Id) {
		response.ErrorWithMessage("同名高级院校已存在！", auc.Ctx)
	}
	if service.IsExistCode(strings.TrimSpace(form.Code), form.Id) {
		response.ErrorWithMessage("同编号高级院校已存在！", auc.Ctx)
	}
	//处理图片上传
	_, _, err := auc.GetFile("badge")
	if err == nil {
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(auc.Ctx, "badge", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), auc.Ctx)
		} else {
			form.Badge = attachmentInfo.Url
		}
	}
	num := service.Update(&form)

	if num > 0 {
		response.Success(auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Del 删除
func (auc *UniversityController) Del() {
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

	var service services.UniversityService
	count := service.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Profile 系统管理-个人资料
func (auc *UniversityController) Profile() {
	auc.Layout = "public/base.html"
	auc.TplName = "admin_user/profile.html"
}

// UpdateNickName 系统管理-个人资料-修改昵称
func (auc *UniversityController) UpdateNickName() {
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
		auc.SetSession(global.LOGIN_USER, *loginAdminUser)
		response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// UpdatePassword 系统管理-个人资料-修改密码
func (auc *UniversityController) UpdatePassword() {
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
func (auc *UniversityController) UpdateAvatar() {
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
			auc.SetSession(global.LOGIN_USER, *loginAdminUser)
			response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, auc.Ctx)
		} else {
			response.Error(auc.Ctx)
		}
	}

}
