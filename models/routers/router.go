package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.WelcomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/main", &controllers.MainController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/change", &controllers.ChangeController{})
	beego.Router("/activities", &controllers.ActivityController{})
	beego.Router("/teacher/main", &controllers.TeacherController{})
	beego.Router("/teacher/add", &controllers.TeacherController{}, "post:Add")
	beego.Router("/teacher/end", &controllers.TeacherController{}, "post:Exit")
	beego.Router("/teacher/accept", &controllers.TeacherController{}, "post:Accept")
	beego.Router("/teacher/addStu", &controllers.TeacherController{}, "post:AddStu")
}
