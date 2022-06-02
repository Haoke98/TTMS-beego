package main

import (
	_ "beego-admin/initialize/conf"
	_ "beego-admin/initialize/mysql"
	_ "beego-admin/initialize/session"
	_ "beego-admin/models"
	_ "beego-admin/routers"
	_ "beego-admin/utils/template"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/context"
	"github.com/beego/beego/v2/client/orm"
	"net/http"
)

func init() {
	var success = []byte("SUPPORT OPTIONS BY SADAM.")
	//跨域设置
	var corsFunc = func(ctx *context.Context) {
		//从请求头拿到Origin写到响应的支持跨域的头部上，这样就可以做到，支持动态跨域。
		origin := ctx.Input.Header("Origin")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", origin)
		//允许访问源
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS,DELETE,PATCH")
		//允许post访问
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,ContentType,Authorization,accept,accept-encoding, authorization, content-type") //header的类型
		ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Input.Method() == http.MethodOptions {
			//所有的options请求通通都返回200
			ctx.Output.SetStatus(http.StatusOK)
			_ = ctx.Output.Body(success)
		}
	}
	beego.InsertFilter("*", beego.BeforeRouter, corsFunc)
}
func main() {

	//输出文件名和行号
	beego.SetLogFuncCall(true)
	orm.RunSyncdb("default", false, true)
	//启动beego
	beego.Run()

}
