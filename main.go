package main

import (
	_ "hello/routers"
	_ "github.com/astaxie/beego/session/mysql"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "gqmms:pf6zbbF2tt@tcp(47.94.91.118:3306)/gqmms?charset=utf8"
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.SetStaticPath("css", "views/css")
	beego.SetStaticPath("image", "views/image")
	beego.SetStaticPath("js", "views/js")
	beego.SetStaticPath("fonts", "views/fonts")
	beego.SetStaticPath("/favicon.ico", "views/image/favicon.ico")
	beego.SetLogger("file", `{"filename":"/www/wwwlogs/gqmms/test.log"}`)
	beego.Run()
}
