package main

import (
	"esee-api/controllers"
	_ "esee-api/routers"
	"esee-api/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	//过滤器
	var FilterUser = func(ctx *context.Context) {
		var Cookie string
		if len(ctx.Request.Header["Cookie"]) != 0 {
			Cookie = ctx.Request.Header["Cookie"][0]
		} else {
			Cookie = ""
		}
		eiamdata := util.GetXtoken(Cookie)
		if eiamdata == "200" {
			requesturl := ctx.Request.RequestURI
			beego.NSNamespace(requesturl,
				beego.NSInclude(
					&controllers.UserController{},
				),
			)
		} else {
			ctx.SetCookie("eiamdata", eiamdata)
			ctx.Redirect(301, "/v1/object/eiam")
		}
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//1. BeforeRouter 寻找路由之前
	//2. BeforeExec 找到路由之后，开始执行相应的 Controller 之前
	//3. AfterExec 执行完 Controller 逻辑之后执行的过滤器
	//4. FinishRouter 执行完逻辑之后执行的过滤器
	// 验证用户是否已经登录
	beego.InsertFilter("/v1/user/*", beego.BeforeRouter, FilterUser)
	beego.Run()

}
