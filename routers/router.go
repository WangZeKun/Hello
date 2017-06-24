// @APIVersion 1.0.0
// @GQMMS WEBAPP API
// @Description web has every tool to get any job done, so codename for the new web APIs.
// @Contact 1015190212@qq.com
package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//网页请求
	//beego.Router("/login", &controllers.HtmlController{}, "get:LoginHtml")
	//beego.Router("/student", &controllers.HtmlController{}, "get:StudentHtml")
	//beego.Router("/teacher", &controllers.HtmlController{}, "get:TeacherHtml")
	//beego.Router("/end", &controllers.HtmlController{}, "get:CollectHtml")
	//beego.Router("/check", &controllers.HtmlController{}, "get:CheckHtml")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.UtilController{},
			),
		),
		beego.NSNamespace("/student",
			beego.NSBefore(func(ctx *context.Context) {
				s := ctx.Input.Session("select")
				if s == nil || s.(string) != "student" {
					ctx.ResponseWriter.WriteHeader(401)
				}
			}),
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),
		beego.NSNamespace("/teacher",
			beego.NSBefore(func(ctx *context.Context) {
				s := ctx.Input.Session("select")
				if s == nil || s.(string) != "teacher" {
					ctx.ResponseWriter.WriteHeader(401)
				}
			}),
			beego.NSInclude(
				&controllers.TeacherController{},
			),
		),
	)
	beego.AddNamespace(ns)
}