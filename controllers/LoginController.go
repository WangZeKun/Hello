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
	if sess != nil {
		c.Redirect("/main", 302)
	}
}

func (c *LoginController) Get() {
	c.TplName = "go.html"
}

func (c *LoginController) Post() {
	user := models.Login{Username: c.GetString("username"), Password: c.GetString("password")}
	b := user.Check()
	if !b {
		c.Data["error"] = "密码错误！"
		c.TplName = "go.html"
	} else {
		c.SetSession("username", user.Username)
		c.Ctx.Redirect(302, "/main")
		//c.Ctx.WriteString(str(err))
	}
}
