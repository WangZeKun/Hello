package controllers

import (
	"hello/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

//教师端API
type TeacherController struct {
	beego.Controller
}

func (c *TeacherController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
}

//@Title 添加活动
//@Description 添加一个新活动
//@Success 200 {string} 成功！
//@Param name formData string true 活动名称
//@Param introduction formData string true 活动简介
//@Param date formData string true 活动时间
//@Param message formData string false 活动的额外信息
//@Failure 500 数据库错误
//@router /add [post]
func (c *TeacherController) Add() {
	message := c.GetString("message")
	beego.Informational(message)
	if message == "" {
		message = "[]"
	}
	activity := models.Activity{
		Name:         c.GetString("name"),
		Introduction: c.GetString("introduction"),
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
	c.Data["json"] = "成功"
	c.ServeJSON()
}

//@Title 结束活动
//@Description 结束一个活动，并写下活动感受。
//@Success 200 {string} 成功！
//@Param id formData int true 活动名称
//@Param score formData string false 活动学分
//@Param impression formData string ture 活动感受
//@Param img formData string false 活动照片(base64)
//@router /end [post]
func (c *TeacherController) Accept() {
	beego.Informational(c.GetString("id"))
	beego.Informational(c.GetString("score"))
	n, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("401")
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
	c.Data["json"] = "成功！"
	c.ServeJSON()
}

//@Title 添加同学
//@Description 为活动添加同学
//Success 200 {string} 成功！或此学生已经报过名了
//Params studentId formData string true 教育ID
//Params activityId formData string true 活动ID
//Failure 500 数据库错误
//@router /addStu [post]
func (c *TeacherController) AddStu() {
	jion := models.Jion{StudentId: c.GetString("studentId"), ActivityId: c.GetString("activityId")}
	b := jion.Check()
	if b {
		c.Data["json"] = "此学生已经报过名了！"
	}else{

		jion.Status = "审核通过"

		err := jion.Insert()
		if err != nil {
			beego.Error("err")
			c.Abort("500")
		}
		c.Data["json"] = "成功！"
	}
	c.ServeJSON()
}

//@Title 得到活动
//@Description返回正在报名的或已经结束的活动
//@Success 200 {object} models.activity.Activity
//@Param status query string true now或end
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /activities [get]
func (c *TeacherController) GetActivties() {
	var activities []models.Activity
	var err error
	if c.GetString("status") == "now" {
		activities, err = models.ShowActivities(c.GetSession("username").(string))
	} else if c.GetString("status") == "end"{
		activities, err = models.ShowAllActivities()
	}else {
		beego.Error("输入错误")
		c.Abort("401")
	}
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["json"] = activities
	c.ServeJSON()
}

//@Title 得到同学
//@Description 返回此活动报名的同学
//@Success 200 {object} models.jion.Jion
//@Param id query int true 活动ID
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /jions [get]
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


//@Title 审核报名
//@Description 通过或不通过学生的报名
//Success 200 {string} 成功！
//Params id formData Array<int> true 审核学生的教育ID
//Params status formData string true 审核状态
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /set [post]
func (c *TeacherController) SetStatus() {
	ids := c.GetString("id")
	var id []int
	err := json.Unmarshal([]byte(ids), id)
	if err != nil {
		beego.Error(err)
		c.Abort("401")
	}
	status := c.GetString("status")
	js := []models.Jion{}
	for i := range id {
		j := models.Jion{Id: i, Status: status}
		js = append(js,j)
	}
	err = models.JionUpdate(js)
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = "成功！"
	c.ServeJSON()
}

//@Title 删除活动
//@Description 删除没有用的活动
//Success 200 {string} 成功！
//@Param id formData int true 删除活动的ID
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /del [post]
func (c *TeacherController) DelActivity() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("401")
	}
	a := models.Activity{Id: id}
	err = a.Delete()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = "成功！"
	c.ServeJSON()
}

//@Title 获得班级
//@Description 返回所在年纪的班级
//@Success 200 {[]string} 班级列表
//@Param grade query string true 年级
//@Failure 500 数据库错误
//@router /getClass [get]
func (c *TeacherController) GetClass() {
	class, err := models.CheckClass(c.GetString("grade"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] =  class
	c.ServeJSON()
}

//@Title 获得同学
//@Description 返回所在班级的同学
//@Success 200 {[]object} models.student.Student
//@Param grade query string true 年级
//@Param class query string true 班级
//@Failure 500 数据库错误
//@router /getStudent [get]
func (c *TeacherController) GetStudent() {
	student, err := models.CheckStudent(c.GetString("grade"), c.GetString("class"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["json"] = student
	c.ServeJSON()
}

//@Title 返回学分
//@Description 返回所在班级/同学的学分
//@Success 200 {[]object} models.model.OutScore
//@Param grade query string true 年级
//@Param class query string true 班级
//@Failure 500 数据库错误
//@router /score [get]
func (c *TeacherController) GetScore() {
	s, err := models.GetScores(c.GetString("class"), c.GetString("grade"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = s
	c.ServeJSON()
}

//@Title 更新活动
//@Description 更新活动信息,并清除掉之前学生们的报名信息
//@Success 200 {[]object} models.model.OutScore
//@Param id dataForm int true 活动ID
//@Param name formData string true 活动名称
//@Param introduction formData string true 活动简介
//@Param date formData string true 活动时间
//@Param message formData Array<string> false 活动的额外信息
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /change [post]
func (c *TeacherController) UdActivity() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Abort("401")
		beego.Error(err)
	}
	activity := models.Activity{
		Id:           id,
	}
	err = activity.Read()
	if err != nil {
		c.Abort("500")
		beego.Error(err)
	}
	if c.GetString("name") != ""{
		activity.Name = c.GetString("name")
	}
	if c.GetString("introduction") != ""{
		activity.Introduction = c.GetString("introduction")
	}
	if c.GetString("date") != ""{
		activity.Date =  c.GetString("date")
	}
	if c.GetString("message") != ""{
		activity.Message = c.GetString("message")
	}
	err = activity.Update()
	if err != nil {
		c.Abort("500")
		beego.Error(err)
	}
	c.Data["json"] = "成功！"
	c.ServeJSON()
}
