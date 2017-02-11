package controllers

import(
	 "github.com/astaxie/beego"
	 "hello/models"
)
type TeacherController struct {
	beego.Controller
}

func (c *TeacherController) Prepare() {
	sess := c.GetSession("username")
	se := c.GetSession("select")
	if sess == nil || se == nil{
		c.Redirect("/login", 302)
	}else if se.(string)=="student"{
		c.Redirect("/main",302)
	}
}
func  (c *TeacherController) Get()  {
    activities, err := models.ShowActivities()
	if err != nil {
		return
	}
	c.Data["activities"] = activities
	c.TplName = "teacher_main.html"
}