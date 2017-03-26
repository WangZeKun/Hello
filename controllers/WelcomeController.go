package controllers

import "github.com/astaxie/beego"

type WelcomeController struct {
	beego.Controller
}

func (c *WelcomeController) Prepare() {
	sess := c.GetSession("username")
	se := c.GetSession("select")
	if sess == nil || se == nil {
		c.Redirect("/login", 302)
	} else if se.(string) == "teacher" {
		c.Redirect("/teacher", 302)
	} else if se.(string) == "student" {
		c.Redirect("/student", 302)
	}
}
