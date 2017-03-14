package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
)

type CollectController struct {
	beego.Controller
}

func (c *CollectController) Prepare() {
	sess := c.GetSession("username")
	se := c.GetSession("select")
	if sess == nil || se == nil {
		c.Redirect("/login", 302)
	} else if sess.(string) != "root" {
		if se.(string) == "student" {
			c.Redirect("/main", 302)
		} else {
			c.Redirect("/teacher/main", 302)
		}
	}
}

func (c *CollectController) Get() {
	c.TplName = "collect.html"
}

func (c *CollectController) Check() {
	c.TplName = "check.html"
}

func (c *CollectController) GetActivtie() {
	a, err := models.ShowAllActivities()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = a
	c.ServeJSON()
}

func (c *CollectController) GetJions() {
	n, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	a := models.Activity{Id: n}
	j, err := a.ShowWhoJion()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = j
	c.ServeJSON()
}

func (c *CollectController) GetScore() {
	s, err := models.GetScores(c.GetString("class"), c.GetString("grade"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = s
	c.ServeJSON()
}

func (c *CollectController) GetClass() {
	class, err := models.CheckClass(c.GetString("grade"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = class
	c.ServeJSON()
}
