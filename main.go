package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "hello/routers"
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
	beego.SetStaticPath("/css", "views/css")
	beego.SetStaticPath("/images", "views/image")
	beego.SetStaticPath("/js", "views/js")
	beego.SetStaticPath("/fonts", "views/fonts")
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.InsertFilter("/teacher", beego.BeforeRouter, FilterTeacher)
	beego.InsertFilter("/student", beego.BeforeRouter, FilterStudent)
	beego.InsertFilter("/message/student/*", beego.BeforeRouter, FilterMessageStudent)
	beego.InsertFilter("/message/teacher/*", beego.BeforeRouter, FilterMessageTeacher)
	beego.Run()
}
