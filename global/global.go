package global

import "TTMS/conf"

// URL_CURRENT 不做任何操作
const URL_CURRENT = "url://current"

// URL_RELOAD 刷新页面
const URL_RELOAD = "url://reload"

// URL_BACK 返回上一个页面
const URL_BACK = "url://back"

// URL_CLOSE_LAYER 关闭当前layer弹窗
const URL_CLOSE_LAYER = "url://close-layer"

// URL_CLOSE_REFRESH 关闭当前弹窗并刷新父级
const URL_CLOSE_REFRESH = "url://close-refresh"

// LOGIN_ADMIN_USER 登录用户key
const LOGIN_ADMIN_USER = "loginAdminUser"

// LOGIN_ADMIN_USER_ID 登录用户id
const LOGIN_ADMIN_USER_ID = "LoginAdminUserId"

// LOGIN_ADMIN_USER_ID_SIGN 登录用户签名
const LOGIN_ADMIN_USER_ID_SIGN = "loginAdminUserIdSign"

const LOGIN_USER = "loginUser"
const LOGIN_USER_ID = "LoginUserId"
const LOGIN_USER_ID_SIGN = "loginUserIdSign"

var (
	// BA_CONFIG conf.Server
	BA_CONFIG conf.Server
)
