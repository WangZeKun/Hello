package controllers

import (
	"encoding/json"
	"hello/models"
	"strconv"

	"github.com/astaxie/beego"
)

type TeacherController struct {
	beego.Controller
}

func (c *TeacherController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
}

func (c *TeacherController) Get() {
	sess := c.GetSession("username")
	teacher := models.Teacher{Id: sess.(string)}
	teacher.Read()
	c.Data["teacher"] = teacher
	c.TplName = "teacher.html"
}

func (c *TeacherController) Add() {
	activity := models.Activity{
		Name:         c.GetString("Name"),
		Introduction: c.GetString("Introduction"),
		WhoBuild:     c.GetSession("username").(string),
		Date:         c.GetString("date"),
	}
	err := activity.Insert()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	//	c.Redirect("/teacher/main", 302)
	c.ServeJSON()
}

func (c *TeacherController) Accept() {
	ob := make(map[string]string)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	n, err := strconv.Atoi(ob["id"])
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
	n, err = strconv.Atoi(ob["score"])
	activity.Score = n
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	activity.Impression = ob["impression"]
	activity.ImagePath = ob["img"]
	err = activity.Update()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.ServeJSON()
}

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
	c.ServeJSON()
}

func (c *TeacherController) GetActivties() {
	activities, err := models.ShowActivities(c.GetSession("username").(string))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["json"] = activities
	c.ServeJSON()
}

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
	c.Data["json"] = j
	c.ServeJSON()
}

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
	c.ServeJSON()
}

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
	c.ServeJSON()
}

func (c *TeacherController) GetClass() {
	class, err := models.CheckClass(c.GetString("grade"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = class
	c.ServeJSON()
}

func (c *TeacherController) GetStudent() {
	student, err := models.CheckStudent(c.GetString("grade"), c.GetString("class"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["json"] = student
	c.ServeJSON()
	return
}
