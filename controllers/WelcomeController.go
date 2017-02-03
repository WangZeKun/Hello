package controllers

import "github.com/astaxie/beego"

type WelcomeController struct {
	beego.Controller
}

func (c *WelcomeController) Prepare() {
	sess := c.GetSession("username")
	if sess != nil {
		//c.Ctx.WriteString("Please login!\n")
		c.Redirect("/mian", 302)
	} else {
		c.Redirect("/login", 302)
	}
}
