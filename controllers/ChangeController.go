package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

type ChangeController struct {
	beego.Controller
}

func (c *ChangeController) Prepare() {
	sess := c.GetSession("username")
	if sess == nil {
		c.Redirect("/login", 302)
	} else {
		c.Layout = "layout.html"
	}
}

func (c *ChangeController) Get() {
	//c.Data["error"] = ""
	c.TplName = "change.tpl"
}

func (c *ChangeController) Post() {
	sess := c.GetSession("username")
	passwordFormer := c.GetString("password_Former")
	password := c.GetString("password")
	passwordRepeat := c.GetString("password_repeat")
	user := models.Login{Username: sess.(string)}
	err := user.Read()
	if err != nil {
		c.Data["error"] = "网络错误，请稍后再试"
		c.TplName = "change.tpl"
	} else if passwordFormer == user.Password {
		c.Data["error"] = "原密码错误"
		c.TplName = "change.tpl"
	} else if password != passwordRepeat {
		c.Data["error"] = "两次输入密码不一致"
		c.TplName = "change.tpl"
	} else {
		user.Password = password
		err = user.Update()
		if err == nil {
			c.Data["Text"] = "修改成功！"
			c.TplName = "message.tpl"
		}
	}
}
