package controllers

import "github.com/astaxie/beego"

type ExitController struct {
	beego.Controller
}

func (c *ExitController) Prepare() {
	sess := c.GetSession("username")
	if sess == nil {
		c.Redirect("/login", 302)
	}
}

func (c *ExitController) Get() {
	c.DelSession("username")
	c.Redirect("/login", 302)
}
