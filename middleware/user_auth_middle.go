package middleware

import (
	"TTMS/global"
	"TTMS/global/response"
	"TTMS/models"
	"TTMS/services"
	"fmt"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/context"
	context2 "github.com/beego/beego/v2/server/web/context"
	"strconv"
	"strings"
)

// UserAuthMiddle 普通用户鉴权中间件
func UserAuthMiddle() {

	//不需要验证的url
	authExcept := map[string]interface{}{
		"client/auth/login":           0,
		"client/auth/check_login":     1,
		"client/auth/logout":          2,
		"client/auth/captcha":         3,
		"client/editor/server":        4,
		"client/auth/refresh_captcha": 5,
	}

	//登录认证中间件过滤器
	var filterLogin = func(ctx *context.Context) {
		url := strings.TrimLeft(ctx.Input.URL(), "/")

		//需要进行登录验证
		if !isUserAuthExceptUrl(strings.ToLower(url), authExcept) {
			//验证是否登录
			_, isLogin := isUserLogin(ctx)
			if !isLogin {
				//ctx.Abort(http.StatusUnauthorized, "请登录")
				response.Result(response.UN_AUTHORIZE, "未登录", "", "", 0, map[string]string{}, (*context2.Context)(ctx))
				return
			}
			//TODO:验证，是否有权限访问
		}

		checkAuth, _ := strconv.Atoi(ctx.Request.PostForm.Get("check_auth"))

		if checkAuth == 1 {
			response.Success((*context2.Context)(ctx))
			return
		}
	}

	beego.InsertFilter("/client/*", beego.BeforeRouter, filterLogin)
}

//判断是否是不需要验证登录的url,只针对admin模块路由的判断
func isUserAuthExceptUrl(url string, m map[string]interface{}) bool {
	urlArr := strings.Split(url, "/")
	if len(urlArr) > 3 {
		url = fmt.Sprintf("%s/%s/%s", urlArr[0], urlArr[1], urlArr[2])
	}
	_, ok := m[url]
	if ok {
		return true
	}
	return false
}

//是否登录
func isUserLogin(ctx *context.Context) (*models.User, bool) {
	loginUser, ok := ctx.Input.Session(global.LOGIN_USER).(models.User)
	if !ok {
		loginUserIDStr := ctx.GetCookie(global.LOGIN_USER_ID)
		loginUserIDSign := ctx.GetCookie(global.LOGIN_USER_ID_SIGN)

		if loginUserIDStr != "" && loginUserIDSign != "" {
			loginUserID, _ := strconv.Atoi(loginUserIDStr)
			var userService services.UserService
			loginUserPointer := userService.GetUserById(loginUserID)

			if loginUserPointer != nil && loginUserPointer.GetSignStr((*context2.Context)(ctx)) == loginUserIDSign {
				ctx.Output.Session(global.LOGIN_USER, *loginUserPointer)
				return loginUserPointer, true
			}
		}
		return nil, false
	}

	return &loginUser, true
}
