package routers

import (
	"beego-admin/controllers"
	"beego-admin/middleware"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dchest/captcha"
	"net/http"
)

func init() {
	//授权登录中间件
	middleware.AuthMiddle()

	web.Get("/", func(ctx *context.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index/index")
	})

	//admin模块路由
	admin := web.NewNamespace("/admin/",
		//UEditor控制器
		web.NSRouter("editor/server", &controllers.EditorController{}, "get,post:Server"),

		//鉴权
		web.NSNamespace("auth/",
			//登录页
			web.NSRouter("login", &controllers.AdminAuthController{}, "get:Login"),
			//退出登录
			web.NSRouter("logout", &controllers.AdminAuthController{}, "get:Logout"),
			//二维码图片输出
			web.NSHandler("captcha/*.png", captcha.Server(240, 80)),
			//登录认证
			web.NSRouter("check_login", &controllers.AdminAuthController{}, "post:CheckLogin"),
			//刷新验证码
			web.NSRouter("refresh_captcha", &controllers.AdminAuthController{}, "post:RefreshCaptcha"),
		),

		//首页
		web.NSRouter("index/index", &controllers.IndexController{}, "get:Index"),

		//操作日志
		web.NSRouter("admin_log/index", &controllers.AdminLogController{}, "get:Index"),

		//菜单管理
		web.NSNamespace("admin_menu/",
			web.NSRouter("index", &controllers.AdminMenuController{}, "get:Index"),
			//菜单管理-添加菜单-界面
			web.NSRouter("add", &controllers.AdminMenuController{}, "get:Add"),
			//菜单管理-添加菜单-创建
			web.NSRouter("create", &controllers.AdminMenuController{}, "post:Create"),
			//菜单管理-修改菜单-界面
			web.NSRouter("edit", &controllers.AdminMenuController{}, "get:Edit"),
			//菜单管理-更新菜单
			web.NSRouter("update", &controllers.AdminMenuController{}, "post:Update"),
			//菜单管理-删除菜单
			web.NSRouter("del", &controllers.AdminMenuController{}, "post:Del"),
		),

		//用户管理
		web.NSNamespace("admin_user/",
			web.NSRouter("index", &controllers.AdminUserController{}, "get:Index"),
			//系统管理-个人资料
			web.NSRouter("profile", &controllers.AdminUserController{}, "get:Profile"),
			//系统管理-个人资料-修改昵称
			web.NSRouter("update_nickname", &controllers.AdminUserController{}, "post:UpdateNickName"),
			//系统管理-个人资料-修改密码
			web.NSRouter("update_password", &controllers.AdminUserController{}, "post:UpdatePassword"),
			//系统管理-个人资料-修改头像
			web.NSRouter("update_avatar", &controllers.AdminUserController{}, "post:UpdateAvatar"),
			//系统管理-用户管理-添加界面
			web.NSRouter("add", &controllers.AdminUserController{}, "get:Add"),
			//系统管理-用户管理-添加
			web.NSRouter("create", &controllers.AdminUserController{}, "post:Create"),
			//系统管理-用户管理-修改界面
			web.NSRouter("edit", &controllers.AdminUserController{}, "get:Edit"),
			//系统管理-用户管理-修改
			web.NSRouter("update", &controllers.AdminUserController{}, "post:Update"),
			//系统管理-用户管理-启用
			web.NSRouter("enable", &controllers.AdminUserController{}, "post:Enable"),
			//系统管理-用户管理-禁用
			web.NSRouter("disable", &controllers.AdminUserController{}, "post:Disable"),
			//系统管理-用户管理-删除
			web.NSRouter("del", &controllers.AdminUserController{}, "post:Del"),
		),

		//系统管理-角色管理
		web.NSNamespace("admin_role/",
			//首页
			web.NSRouter("index", &controllers.AdminRoleController{}, "get:Index"),
			//系统管理-角色管理-添加界面
			web.NSRouter("add", &controllers.AdminRoleController{}, "get:Add"),
			//系统管理-角色管理-添加
			web.NSRouter("create", &controllers.AdminRoleController{}, "post:Create"),
			//菜单管理-角色管理-修改界面
			web.NSRouter("edit", &controllers.AdminRoleController{}, "get:Edit"),
			//菜单管理-角色管理-修改
			web.NSRouter("update", &controllers.AdminRoleController{}, "post:Update"),
			//菜单管理-角色管理-删除
			web.NSRouter("del", &controllers.AdminRoleController{}, "post:Del"),
			//菜单管理-角色管理-启用角色
			web.NSRouter("enable", &controllers.AdminRoleController{}, "post:Enable"),
			//菜单管理-角色管理-禁用角色
			web.NSRouter("disable", &controllers.AdminRoleController{}, "post:Disable"),
			//菜单管理-角色管理-角色授权界面
			web.NSRouter("access", &controllers.AdminRoleController{}, "get:Access"),
			//菜单管理-角色管理-角色授权
			web.NSRouter("access_operate", &controllers.AdminRoleController{}, "post:AccessOperate"),
		),
		//设置中心
		web.NSNamespace("setting/",
			//设置中心-后台设置
			web.NSRouter("admin", &controllers.SettingController{}, "get:Admin"),
			//设置中心-更新设置
			web.NSRouter("update", &controllers.SettingController{}, "post:Update"),
		),

		//系统管理-开发管理
		web.NSNamespace("database/",
			//系统管理-开发管理-数据维护
			web.NSRouter("table", &controllers.DatabaseController{}, "get:Table"),
			//系统管理-开发管理-数据维护-优化表
			web.NSRouter("optimize", &controllers.DatabaseController{}, "post:Optimize"),
			//系统管理-开发管理-数据维护-修复表
			web.NSRouter("repair", &controllers.DatabaseController{}, "post:Repair"),
			//系统管理-开发管理-数据维护-查看详情
			web.NSRouter("view", &controllers.DatabaseController{}, "get,post:View"),
		),

		//用户等级管理
		web.NSNamespace("user_level/",
			//首页
			web.NSRouter("index", &controllers.UserLevelController{}, "get:Index"),
			//用户等级管理-添加界面
			web.NSRouter("add", &controllers.UserLevelController{}, "get:Add"),
			//用户等级管理-添加
			web.NSRouter("create", &controllers.UserLevelController{}, "post:Create"),
			//用户等级管理-修改界面
			web.NSRouter("edit", &controllers.UserLevelController{}, "get:Edit"),
			//用户等级管理-修改
			web.NSRouter("update", &controllers.UserLevelController{}, "post:Update"),
			//用户等级管理-启用
			web.NSRouter("enable", &controllers.UserLevelController{}, "post:Enable"),
			//用户等级管理-禁用
			web.NSRouter("disable", &controllers.UserLevelController{}, "post:Disable"),
			//用户等级管理-删除
			web.NSRouter("del", &controllers.UserLevelController{}, "post:Del"),
			//用户等级管理-导出
			web.NSRouter("export", &controllers.UserLevelController{}, "get:Export"),
		),

		//用户管理
		web.NSNamespace("user/",
			web.NSRouter("index", &controllers.UserController{}, "get:Index"),
			//用户管理-添加界面
			web.NSRouter("add", &controllers.UserController{}, "get:Add"),
			//用户管理-添加
			web.NSRouter("create", &controllers.UserController{}, "post:Create"),
			//用户管理-修改界面
			web.NSRouter("edit", &controllers.UserController{}, "get:Edit"),
			//用户管理-修改
			web.NSRouter("update", &controllers.UserController{}, "post:Update"),
			//用户管理-启用
			web.NSRouter("enable", &controllers.UserController{}, "post:Enable"),
			//用户管理-禁用
			web.NSRouter("disable", &controllers.UserController{}, "post:Disable"),
			//用户管理-删除
			web.NSRouter("del", &controllers.UserController{}, "post:Del"),
			//用户管理-导出
			web.NSRouter("export", &controllers.UserController{}, "get:Export"),
		),

		web.NSNamespace("train-management/",
			web.NSRouter("index", &controllers.TrainPlanController{}, "get:Index"),
			web.NSRouter("add", &controllers.TrainPlanController{}, "get:Add"),
			web.NSRouter("create", &controllers.TrainPlanController{}, "post:Create"),
			web.NSRouter("edit", &controllers.TrainPlanController{}, "get:Edit"),
			web.NSRouter("update", &controllers.TrainPlanController{}, "post:Update"),
		),

		web.NSNamespace("train-course/",
			web.NSRouter("index", &controllers.TrainCourseController{}, "get:Index"),
			web.NSRouter("add", &controllers.TrainCourseController{}, "get:Add"),
			web.NSRouter("create", &controllers.TrainCourseController{}, "post:Create"),
			web.NSRouter("edit", &controllers.TrainCourseController{}, "get:Edit"),
			web.NSRouter("update", &controllers.TrainCourseController{}, "post:Update"),
		),

		web.NSNamespace("university/",
			web.NSRouter("index", &controllers.UniversityController{}, "get:Index"),
			web.NSRouter("add", &controllers.UniversityController{}, "get:Add"),
			web.NSRouter("create", &controllers.UniversityController{}, "post:Create"),
			web.NSRouter("edit", &controllers.UniversityController{}, "get:Edit"),
			web.NSRouter("update", &controllers.UniversityController{}, "post:Update"),
			web.NSRouter("del", &controllers.UniversityController{}, "post:Del"),
		),

		web.NSNamespace("quota/",
			web.NSRouter("index", &controllers.QuotaController{}, "get:Index"),
			web.NSRouter("update", &controllers.QuotaController{}, "post:Update"),
		),

		web.NSRouter("/petition/index", &controllers.PetitionController{}, "get:Index"),
	)
	train := web.NewNamespace("/train",
		web.NSRouter("/plan", &controllers.TrainPlanController{}, "get:Get"),
	)

	//鉴权
	auth := web.NewNamespace("/auth",
		//登录认证
		web.NSRouter("/login", &controllers.UserAuthController{}, "post:Login"),
		//退出登录
		web.NSRouter("/logout", &controllers.UserAuthController{}, "get:Logout"),
		//二维码图片输出
		web.NSHandler("/captcha/*.png", captcha.Server(240, 80)),
		//刷新验证码
		web.NSRouter("/refresh_captcha", &controllers.UserAuthController{}, "post:RefreshCaptcha"),
	)

	web.AddNamespace(admin, train, auth)
}
