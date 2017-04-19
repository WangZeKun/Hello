package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
	"time"
)

type StudentController struct {
	beego.Controller
}

func (c *StudentController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8100")
}

//获取学生信息
func (c *StudentController) GetStudentMessage() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	err := stu.Read()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["json"] = sendMessage("成功！", stu)
	c.ServeJSON()
}

//获取学生参加活动的信息
func (c *StudentController) GetCanjia() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	j, err := stu.ShowWhatJion()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", j)
	c.ServeJSON()
}

//获得学校发布的信息
func (c *StudentController) GetActivity() {
	name := c.GetString("who")
	if name != "root" {
		var err error
		sess := c.GetSession("username")
		stu := models.Student{Id: sess.(string)}
		name, err = stu.CheckClassTeacher()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
	}
	data, err := models.ShowActivities(name)
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", data)
	c.ServeJSON()
}

//报名活动
func (c *StudentController) SetJion() {
	sess := c.GetSession("username")
	jion := models.Jion{
		ActivityId: c.GetString("activity_id"),
		StudentId:  sess.(string),
		Message:    c.GetString("message"),
	}
	if jion.Message == "" {
		jion.Message = "[]"
	}
	b := jion.Check()
	if b {
		c.Data["json"] = sendMessage("您已经报过名了！", nil)
	} else {
		jion.Status = "审核中"
		jion.Date = time.Now()
		err := jion.Insert()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
		c.Data["json"] = sendMessage("报名成功！", nil)
	}
	c.ServeJSON()
}
