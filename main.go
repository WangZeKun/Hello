package main

import (
	_ "hello/routers"
	_ "github.com/astaxie/beego/session/mysql"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "gqmms:pf6zbbF2tt@tcp(127.0.0.1:3306)/gqmms?charset=utf8"
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.BeeLogger.DelLogger("console")
	beego.SetLevel(beego.LevelError)
	beego.SetLogger("file", `{"filename":"/www/wwwlogs/gqmms/test.log"}`)
	beego.Run()
}
