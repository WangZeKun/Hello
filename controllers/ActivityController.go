package controllers

import (
	"hello/models"

	"time"

	"github.com/astaxie/beego"
)

type ActivityController struct {
	beego.Controller
}

func (c *ActivityController) Prepare() {
	sess := c.GetSession("username")
	if sess == nil {
		c.Redirect("/login", 302)
	} else {
		c.Layout = "layout.html"
	}
}

func (c *ActivityController) Get() {
	activities, err := models.ShowActivities()
	if err != nil {
		return
	}
	c.Data["activities"] = activities
	c.TplName = "activity.tpl"
}

func (c *ActivityController) Post() {
	sess := c.GetSession("username")
	id := c.GetString("activity_id")
	activity := models.Activity{Id:id}
	err := activity.Read()
	if err != nil{
		beego.Error(err)
	}
	m := make(map[string] string)  
	for _,i := range activity.GetMessage(){
		m[i]=c.GetString(i)
	}
	jion := models.Jion{ActivityId: id, StudentId: sess.(string)}
	jion.SetMessage(m)
	b := jion.Check()
	if b {
		c.Data["Text"] = "您已经报名了！"
		c.TplName = "message.tpl"
	} else {
		jion.Status = "审核中"
		jion.Date = time.Now()
		err := jion.Insert()
		if err != nil {
			c.Data["Text"] = "错误，请稍后重试"
			c.TplName = "message.tpl"
		} else {
			c.Data["Text"] = "报名成功"
			c.TplName = "message.tpl"
		}
	}
}
