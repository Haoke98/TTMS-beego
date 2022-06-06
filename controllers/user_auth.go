package controllers

import (
	"TTMS/formvalidate"
	"TTMS/global"
	"TTMS/global/response"
	"TTMS/services"
	"TTMS/utils"
	"encoding/json"
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/dchest/captcha"
	"github.com/gookit/validate"
	"strconv"
)

var userLogService services.UserLogService

// UserAuthController struct
type UserAuthController struct {
	baseController
}

// Login 普通用户登录认证
func (uac *UserAuthController) Login() {
	//数据校验
	valid := validation.Validation{}
	loginForm := formvalidate.LoginForm{}
	if err := json.Unmarshal(uac.Ctx.Input.RequestBody, &loginForm); err != nil {
		response.ErrorWithMessage(err.Error(), uac.Ctx)
	}
	//
	//if err := uac.ParseForm(&loginForm); err != nil {
	//	response.ErrorWithMessage(err.Error(), uac.Ctx)
	//}

	v := validate.Struct(loginForm)

	//TODO：这里图形验证码是必须的，但是初期开发就先不高的那么复杂，不进行验证。
	//看是否需要校验验证码
	isCaptcha, _ := strconv.Atoi(global.BA_CONFIG.Login.Captcha)
	if isCaptcha > 0 {
		valid.Required(loginForm.Captcha, "captcha").Message("请输入验证码.")
		if ok := captcha.VerifyString(loginForm.CaptchaId, loginForm.Captcha); !ok {
			response.ErrorWithMessage("验证码错误.", uac.Ctx)
		}
	}

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uac.Ctx)
	}

	//基础验证通过后，进行用户验证
	var userService services.UserService
	loginUser, err := userService.CheckLogin(loginForm, uac.Ctx)
	if err != nil {
		response.ErrorWithMessage(err.Error(), uac.Ctx)
	}

	//登录日志记录
	userLogService.LoginLog(loginUser.Id, uac.Ctx)

	redirect, _ := uac.GetSession("redirect").(string)
	if redirect != "" {
		response.SuccessWithMessageAndUrl("登录成功", redirect, uac.Ctx)
	} else {
		response.SuccessWithMessageAndUrl("登录成功", "/admin/index/index", uac.Ctx)
	}
}

// Logout 退出登录
func (uac *UserAuthController) Logout() {
	uac.DelSession(global.LOGIN_USER)
	uac.Ctx.SetCookie(global.LOGIN_USER_ID, "", -1)
	uac.Ctx.SetCookie(global.LOGIN_USER_ID_SIGN, "", -1)
}

// RefreshCaptcha 刷新验证码
func (uac *UserAuthController) RefreshCaptcha() {
	captchaID := uac.GetString("captchaId")
	res := map[string]interface{}{
		"isNew": false,
	}
	if captchaID == "" {
		res["msg"] = "参数错误."
	}

	isReload := captcha.Reload(captchaID)
	if isReload {
		res["captchaId"] = captchaID
	} else {
		res["isNew"] = true
		res["captcha"] = utils.GetCaptcha()
	}

	uac.Data["json"] = res

	uac.ServeJSON()
}
