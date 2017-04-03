package controllers

import (
	"time"

	"hello/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
}

func (c *MainController) Get() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	err := stu.Read()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["stu"] = &stu
	c.TplName = "student.html"
}

func (c *MainController) GetCanjia() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	j, err := stu.ShowWhatJion()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = models.SendMessage("成功！", j)
	c.ServeJSON()
}

func (c *MainController) GetRootActivity() {
	data, err := models.ShowActivities("root")
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = models.SendMessage("成功！", data)
	c.ServeJSON()
}

func (c *MainController) GetClassActivity() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	class, err := stu.CheckClass()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	data, err := models.ShowActivities(class)
	c.Data["json"] = models.SendMessage("成功！", data)
	c.ServeJSON()
}

func (c *MainController) GetGradeActivity() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	grade, err := stu.CheckGrade()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	data, err := models.ShowActivities(grade)
	c.Data["json"] = models.SendMessage("成功！", data)
	c.ServeJSON()
}

func (c *MainController) SetJion() {
	sess := c.GetSession("username")
	id, err := c.GetInt("activity_id")
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	activity := models.Activity{Id: id}
	err = activity.Read()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	jion := models.Jion{ActivityId: c.GetString("activity_id"), StudentId: sess.(string)}
	b := jion.Check()
	if b {
		c.Data["json"] = models.SendMessage("您已经报过名了！", nil)
	} else {
		jion.Status = "审核中"
		jion.Date = time.Now()
		err := jion.Insert()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
		c.Data["json"] = models.SendMessage("报名成功！", nil)
	}
	c.ServeJSON()
}

func (c *MainController) Change() {
	sess := c.GetSession("username")
	user := models.Login{Username: sess.(string)}
	err := user.Read()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	if user.Password != c.GetString("password") {
		c.Data["json"] = "原密码错误！"
	} else {
		user.Password = c.GetString("newpassword")
		err = user.Update()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
		c.Data["json"] = models.SendMessage("修改成功", nil)
	}
	c.ServeJSON()
}
