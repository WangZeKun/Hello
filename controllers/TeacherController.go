package controllers

import (
	"hello/models"
	"github.com/astaxie/beego"
)

type TeacherController struct {
	beego.Controller
}

func (c *TeacherController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8100")
}

//添加活动
func (c *TeacherController) Add() {
	message := c.GetString("message")
	beego.Informational(message)
	if message == "" {
		message = "[]"
	}
	activity := models.Activity{
		Name:         c.GetString("Name"),
		Introduction: c.GetString("Introduction"),
		WhoBuild:     c.GetSession("username").(string),
		Date:         c.GetString("date"),
		Message:      message,
	}
	beego.Informational(activity.Message)
	err := activity.Insert()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", nil)
	c.ServeJSON()
}

//结束活动
func (c *TeacherController) Accept() {
	beego.Informational(c.GetString("id"))
	beego.Informational(c.GetString("score"))
	n, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	activity := models.Activity{Id: n}
	err = activity.Read()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	n, err = c.GetInt("score")
	activity.Score = n
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	activity.Impression = c.GetString("impression")
	activity.ImagePath = c.GetString("img")
	err = activity.EndActivity()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", nil)
	c.ServeJSON()
}

//添加同学
func (c *TeacherController) AddStu() {
	jion := models.Jion{StudentId: c.GetString("name"), ActivityId: c.GetString("activityId")}
	b := jion.Check()
	if b {
		return
	}
	jion.Status = "审核通过"

	err := jion.Insert()
	if err != nil {
		beego.Error("err")
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", nil)
	c.ServeJSON()
}

//返回正在报名的活动
func (c *TeacherController) GetActivties() {
	var activities []models.Activity
	var err error
	if c.GetString("status") == "now"{
		activities, err = models.ShowActivities(c.GetSession("username").(string))
	}else {
		activities, err = models.ShowAllActivities()
	}
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["json"] = sendMessage("成功！", activities)
	c.ServeJSON()
}

//返回此活动报名的同学
func (c *TeacherController) GetJions() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	a := models.Activity{Id: id}
	j, err := a.ShowWhoJion()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", j)
	c.ServeJSON()
}

//通过或不通过学生的报名
func (c *TeacherController) SetStatus() {
	id, err := c.GetInt("jionid")
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	status := c.GetString("status")
	j := models.Jion{Id: id, Status: status}
	err = j.Update()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", nil)
	c.ServeJSON()
}

//删除活动
func (c *TeacherController) DelActivity() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	a := models.Activity{Id: id}
	err = a.Delete()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", nil)
	c.ServeJSON()
}

//返回所在年纪的班级
func (c *TeacherController) GetClass() {
	class, err := models.CheckClass(c.GetString("grade"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", class)
	c.ServeJSON()
}

//返回所在班级的同学
func (c *TeacherController) GetStudent() {
	student, err := models.CheckStudent(c.GetString("grade"), c.GetString("class"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["json"] = sendMessage("成功！", student)
	c.ServeJSON()
	return
}

//返回学分
func (c *TeacherController) GetScore() {
	s, err := models.GetScores(c.GetString("class"), c.GetString("grade"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = sendMessage("成功！", s)
	c.ServeJSON()
}

//更新活动
func (c *TeacherController) UdActivity(){
	id,err := c.GetInt("id")
	if err != nil{
		c.Abort("500")
		beego.Error(err)
	}
	activity := models.Activity{
		Id:id,
		Date:c.GetString("date"),
		Introduction:c.GetString("introduction"),
		Message:c.GetString("message"),
	}
	err = activity.Update()
	if err != nil {
		c.Abort("500")
		beego.Error(err)
	}
	return
}
