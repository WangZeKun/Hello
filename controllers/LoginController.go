package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	user := models.Login{Username: c.GetString("username")}
	b := user.Check(c.GetString("password"))
	if !b {
		c.Data["json"] = sendMessage("密码错误！", nil)
		c.ServeJSON()
	} else {
		if c.GetString("select") == "学生登陆" && user.Who == "student" {
			c.SetSession("username", user.Username)
			c.SetSession("select", user.Who)
			sess:= c.StartSession()
			beego.Informational(sess.SessionID())
			c.Data["json"] = sendMessage("成功，学生登录！", sess.SessionID())
		} else if c.GetString("select") == "教师登陆" && user.Who == "teacher" {
			c.Data["json"] = sendMessage("成功，教师登录！", nil)
			c.SetSession("username", user.Username)
			c.SetSession("select", user.Who)
		} else {
			c.Data["json"] = sendMessage("请选择正确的登录用户！", nil)
		}

		c.ServeJSON()
		//c.Ctx.WriteString(str(err))
	}
}
