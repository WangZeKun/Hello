package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
)

type HtmlController struct {
	beego.Controller
}

func (c *HtmlController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
}

func (c *HtmlController) LoginHtml() {
	c.TplName = "login.html"
}

func (c *HtmlController) CollectHtml() {
	c.TplName = "collect.html"
}

func (c *HtmlController) CheckHtml() {
	c.TplName = "check.html"
}

func (c *HtmlController) TeacherHtml() {
	sess := c.GetSession("username")
	teacher := models.Teacher{Id: sess.(string)}
	teacher.Read()
	c.Data["teacher"] = teacher
	c.TplName = "teacher.html"
}

func (c *HtmlController) StudentHtml() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	err := stu.Read()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["stu"] = &stu
	c.TplName = "student.html"
}