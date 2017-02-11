package controllers

import "github.com/astaxie/beego"

type ExitController struct {
	beego.Controller
}

func (c *ExitController) Prepare() {
	sess := c.GetSession("username")
	se := c.GetSession("select")
	if sess == nil || se == nil{
		c.Redirect("/login", 302)
	}else if se.(string)=="teacher"{
		c.Redirect("/teacher/main",302)
	} 
}

func (c *ExitController) Get() {
	c.DelSession("username")
	c.Redirect("/login", 302)
}
