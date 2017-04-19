package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//网页请求
	beego.Router("/login", &controllers.HtmlController{}, "get:LoginHtml")
	beego.Router("/student", &controllers.HtmlController{}, "get:StudentHtml")
	beego.Router("/teacher", &controllers.HtmlController{}, "get:TeacherHtml")
	beego.Router("/end", &controllers.HtmlController{}, "get:CollectHtml")
	beego.Router("/check", &controllers.HtmlController{}, "get:CheckHtml")

	//一些工具类的请求
	beego.Router("/", &controllers.UtilController{}, "get:JumpTo")
	beego.Router("/message/text", &controllers.UtilController{}, "get:Text")
	beego.Router("/message/exit", &controllers.UtilController{}, "get:Exit")
	beego.Router("/message/change", &controllers.UtilController{}, "get:Change")
	beego.Router("/message/login", &controllers.UtilController{}, "get:Login")

	//学生界面请求
	beego.Router("/message/student", &controllers.StudentController{}, "get:GetStudentMessage")
	beego.Router("/message/student/canjia", &controllers.StudentController{}, "get:GetCanjia")
	beego.Router("/message/student/activity", &controllers.StudentController{}, "get:GetActivity")
	beego.Router("/message/student/jion", &controllers.StudentController{}, "get:SetJion")

	//教师界面请求
	beego.Router("/message/teacher/add", &controllers.TeacherController{}, "get,post:Add")
	beego.Router("/message/teacher/accept", &controllers.TeacherController{}, "post:Accept")
	beego.Router("/message/teacher/addStu", &controllers.TeacherController{}, "get:AddStu")
	beego.Router("/message/teacher/activities", &controllers.TeacherController{}, "get:GetActivties")
	beego.Router("/message/teacher/activity", &controllers.TeacherController{}, "get:GetJions")
	beego.Router("/message/teacher/set", &controllers.TeacherController{}, "get:SetStatus")
	beego.Router("/message/teacher/del", &controllers.TeacherController{}, "get:DelActivity")
	beego.Router("/message/teacher/getclass", &controllers.TeacherController{}, "get:GetClass")
	beego.Router("/message/teacher/getstudent", &controllers.TeacherController{}, "get:GetStudent")
	beego.Router("/message/teacher/score", &controllers.TeacherController{}, "get:GetScore")
	beego.Router("/message/teacher/change", &controllers.TeacherController{},"get:UdActivity")
}
