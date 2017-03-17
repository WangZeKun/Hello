package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.WelcomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/message/login", &controllers.LoginController{}, "get:Post")
	beego.Router("/student", &controllers.MainController{})
	beego.Router("/message/student/canjia", &controllers.MainController{}, "get:GetCanjia")
	beego.Router("/message/student/root", &controllers.MainController{}, "get:GetRootActivity")
	beego.Router("/message/student/class", &controllers.MainController{}, "get:GetClassActivity")
	beego.Router("/message/student/grade", &controllers.MainController{}, "get:GetGradeActivity")
	beego.Router("/message/student/change", &controllers.MainController{}, "get:Change")
	beego.Router("/message/student/jion", &controllers.MainController{}, "get:SetJion")
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/teacher", &controllers.TeacherController{})
	beego.Router("/message/teacher/add", &controllers.TeacherController{}, "get:Add")
	beego.Router("/message/teacher/accept", &controllers.TeacherController{}, "post:Accept")
	beego.Router("/message/teacher/addStu", &controllers.TeacherController{}, "get:AddStu")
	beego.Router("/message/teacher/activities", &controllers.TeacherController{}, "get:GetActivties")
	beego.Router("/message/teacher/activity", &controllers.TeacherController{}, "get:GetJions")
	beego.Router("/message/teacher/set", &controllers.TeacherController{}, "get:SetStatus")
	beego.Router("/message/teacher/del", &controllers.TeacherController{}, "get:DelActivity")
	beego.Router("/message/teacher/getclass", &controllers.TeacherController{}, "get:GetClass")
	beego.Router("/message/teacher/getstudent", &controllers.TeacherController{}, "get:GetStudent")
	beego.Router("/message/teacher/Cactivities", &controllers.CollectController{}, "get:GetActivtie")
	beego.Router("/message/teacher/Cactivity", &controllers.CollectController{}, "get:GetJions")
	beego.Router("/end", &controllers.CollectController{})
	beego.Router("/check", &controllers.CollectController{}, "get:Check")
	beego.Router("/message/teacher/score", &controllers.CollectController{}, "get:GetScore")
}
