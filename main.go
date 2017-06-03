package main

import (
	_ "hello/routers"
	_ "github.com/astaxie/beego/session/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)


var FilterTeacher = func(ctx *context.Context) {
	s := ctx.Input.Session("select")
	if s == nil {
		ctx.Redirect(302, "/login")
	} else if s == "student" {
		ctx.Redirect(302, "/student")
	}
}

var FilterStudent = func(ctx *context.Context) {
	s := ctx.Input.Session("select")
	if s == nil {
		ctx.Redirect(302, "/login")
	} else if s.(string) == "teacher" {
		ctx.Redirect(302, "/teacher")
	}
}

var FilterMessageStudent = func(ctx *context.Context) {
	s := ctx.Input.Session("select")
	if s == nil || s.(string) == "teacher" {
		ctx.Abort(401, "")
	}
}

var FilterMessageTeacher = func(ctx *context.Context) {
	s := ctx.Input.Session("select")
	if s == nil || s.(string) == "student" {
		ctx.Abort(401, "")
	}
}

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "gqmms:pf6zbbF2tt@tcp(127.0.0.1:3306)/gqmms?charset=utf8"
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.SetStaticPath("css", "views/css")
	beego.SetStaticPath("image", "views/image")
	beego.SetStaticPath("js", "views/js")
	beego.SetStaticPath("fonts", "views/fonts")
	beego.SetStaticPath("/favicon.ico", "views/image/favicon.ico")
	beego.SetLogger("file", `{"filename":"/www/wwwlogs/gqmms/test.log"}`)
	beego.BeeLogger.DelLogger("console")
	beego.InsertFilter("/teacher", beego.BeforeRouter, FilterTeacher)
	beego.InsertFilter("/student", beego.BeforeRouter, FilterStudent)
	beego.InsertFilter("/message/student/*", beego.BeforeRouter, FilterMessageStudent)
	beego.InsertFilter("/message/teacher/*", beego.BeforeRouter, FilterMessageTeacher)
	beego.Run()
}
