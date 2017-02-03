package controllers

import "github.com/astaxie/beego"

type TeacherController struct {
	beego.Controller
}

func  (c *TeacherController) Get()  {
    c.Layout = "layout.html"
    c.TplName = "activity_set.tpl"
}