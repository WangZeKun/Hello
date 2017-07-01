package controllers

import (
	"encoding/json"
	"hello/models"

	"github.com/astaxie/beego"
)

//教师端API
type TeacherController struct {
	beego.Controller
}

func (c *TeacherController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "https://www.gqmms.wang")
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
	activity := models.Activity{
		Name:         c.GetString("name"),
		Introduction: c.GetString("introduction"),
		WhoBuild:     c.GetSession("username").(string),
		Date:         c.GetString("date"),
	}
	err := json.Unmarshal([]byte(c.GetString("message")), &activity.Message)
	beego.Informational(activity.Message)
	err = activity.Insert()
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
//@Param impression formData string true 活动感受
//@Param img formData string false 活动照片(base64)
//@router /end [post]
func (c *TeacherController) Accept() {
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
	score, err := c.GetInt("score")
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	activity.Impression = c.GetString("impression")
	var imgs []string
	json.Unmarshal([]byte(c.GetString("img")), &imgs)
	for _, img := range imgs {
		i := models.Photo{
			ActivityId: n,
			Photo:      img,
		}
		err = i.Insert()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
	}
	err = activity.EndActivity(score)
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = "成功！"
	c.ServeJSON()
}

//@Title 添加同学
//@Description 为活动添加同学
//@Success 200 {string} 成功！或此学生已经报过名了
//@Params studentId formData string true 教育ID
//@Params activityId formData string true 活动ID
//@Failure 500 数据库错误
//@router /addStu [post]
func (c *TeacherController) AddStu() {
	activityId, _ := c.GetInt("activityId")
	join := models.Join{StudentId: c.GetString("studentId"), ActivityId: activityId}
	b, err := join.Check()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	if b {
		c.Data["json"] = "此学生已经报过名了！"
	} else {

		join.Status = "审核通过"

		err := join.Insert()
		if err != nil {
			beego.Error(err)
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
		activities, err = models.ShowActivities(c.GetSession("username").(string), true)
	} else if c.GetString("status") == "end" {
		activities, err = models.ShowActivities(c.GetSession("username").(string), false)
	} else {
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
//@Success 200 {object} models.join.Join
//@Param id query int true 活动ID
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /joins [get]
func (c *TeacherController) Getjoins() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	a := models.Activity{Id: id}
	j, err := a.ShowWhoJoin()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = j
	c.ServeJSON()
}

//@Title 审核报名
//@Description 通过或不通过学生的报名
//@Success 200 {string} 成功！
//@Params id formData Array<int> true 审核学生的教育ID
//@Params status formData string true 审核状态
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /set [post]
func (c *TeacherController) SetStatus() {
	ids := "[" + c.GetString("id") + "]"
	beego.Informational(ids)
	var id []int
	err := json.Unmarshal([]byte(ids), &id)
	beego.Informational(id)
	if err != nil {
		beego.Error(err)
		c.Abort("401")
	}
	status := c.GetString("status")
	for _, i := range id {
		j := models.Join{Id: i, Status: status}
		err = j.Update()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
	}
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
	c.Data["json"] = class
	c.ServeJSON()
}

//@Title 获得同学
//@Description 返回所在班级的同学
//@Success 200 {[]object} models.student.Number
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
//@Success 200 {object} models.model.OutScore
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
		c.Abort("400")
		beego.Error(err)
	}
	activity := models.Activity{
		Id:           id,
		Name:         c.GetString("name"),
		Introduction: c.GetString("introduction"),
		Date:         c.GetString("date"),
	}
	err = json.Unmarshal([]byte(c.GetString("message")), &activity.Message)
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	err = activity.Update()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = "成功！"
	c.ServeJSON()
}

//@Title 添加图片
//@Description 添加活动图片
//@Success 200 {object} models.photo.Photo
//@Param id dataForm int true 活动ID
//@Param photo formData string true 图片数据
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /photo [post]
func (c *TeacherController) Photo() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	photo := models.Photo{ActivityId: id, Photo: c.GetString("photo")}
	err = photo.Insert()
}

//@Title 获得图片
//Description 得到活动图片
//Success 200 {object} models.photo.Photo
//Param id query int true 活动ID
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /getPhotos [get]
func (c *TeacherController) GetPhotos() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	photos, err := models.GetPhotos(id)
	c.Data["json"] = photos
	c.ServeJSON()
}
