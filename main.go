package main

import (
	_ "hello/routers"
	_ "github.com/astaxie/beego/session/mysql"

	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "gqmms:pf6zbbF2tt@tcp(127.0.0.1:3306)/gqmms?charset=utf8"
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.SetStaticPath("css", "views/css")
	beego.SetStaticPath("image", "views/image")
	beego.SetStaticPath("js", "views/js")
	beego.SetStaticPath("fonts", "views/fonts")
	beego.SetStaticPath("/favicon.ico", "views/image/favicon.ico")
	beego.SetLogger("file", `{"filename":"/www/wwwlogs/gqmms/test.log"}`)
	beego.Run()
}
