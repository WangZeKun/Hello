package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {
	sess := c.GetSession("username")
	se := c.GetSession("select")
	if sess != nil && se != nil {
		if se.(string) == "student" {
			c.Redirect("/main", 302)
		} else if se.(string) == "teacher" {
			c.Redirect("/teacher/main", 302)
		}
	}
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	user := models.Login{Username: c.GetString("username")}
	b := user.Check(c.GetString("password"))
	if !b {
		c.Data["json"] = "密码错误！"
		c.ServeJSON()
	} else {
		if c.GetString("select") == "学生登陆" && user.Who == "student" {
			c.Data["json"] = "success student"
			c.SetSession("username", user.Username)
			c.SetSession("select", user.Who)
		} else if c.GetString("select") == "教师登陆" && user.Who == "teacher" {
			c.Data["json"] = "success teacher"
			c.SetSession("username", user.Username)
			c.SetSession("select", user.Who)
		} else {
			c.Data["json"] = "pleace choose user type"
		}
		c.ServeJSON()
		//c.Ctx.WriteString(str(err))
	}
}
