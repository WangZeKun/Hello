package controllers

import (
	"fmt"
	"hello/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	sess := c.GetSession("username")
	se := c.GetSession("select")
	if sess == nil || se == nil{
		c.Redirect("/login", 302)
	}else if se.(string)=="teacher"{
		c.Redirect("/teacher/main",302)
	} else {
		c.Layout = "layout.html"
	}
}

func (c *MainController) Get() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	err := stu.Read()
	if err == nil {
		c.Data["stu"] = &stu
		can, err := stu.ShowJion()
		if err == nil {
			c.Data["Canjia"] = can
		}
	}
	c.Layout = "layout.html"
	c.TplName = "main.tpl"
}

func (c *MainController) Post() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	err := stu.Read()
	if err != nil {
		c.Data["Text"] = "网络错误，请稍后再试！"
		c.TplName = "message.tpl"
		return
	}
	telephone := c.GetString("telephone")
	fmt.Println(telephone)
	if telephone != "" {
		stu.Telephone = telephone
	}
	Qq := c.GetString("QQ")
	fmt.Println(Qq)
	if Qq != "" {
		stu.Qq = Qq
	}
	WeChat := c.GetString("weChat")
	fmt.Println(WeChat)
	if WeChat != "" {
		stu.WeChat = WeChat
	}
	Jianjie := c.GetString("jianjie")
	fmt.Println(Jianjie)
	if Jianjie != "" {
		stu.Jianjie = Jianjie
	}
	err = stu.Update()
	if err == nil {
		c.Data["Text"] = "修改成功"
		c.TplName = "message.tpl"
	}
}
