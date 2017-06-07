package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
	"time"
)

//学生端API
type StudentController struct {
	beego.Controller
}

func (c *StudentController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
}


//@Title 得到报名活动信息
//@Description 获取学生参加活动的信息
//@Success 200 {object} models.model.OutStudentJion
//@Failure 500 数据库错误
//@router /canjia [get]
func (c *StudentController) GetCanjia() {
	sess := c.GetSession("username")
	stu := models.Student{Id: sess.(string)}
	j, err := stu.ShowWhatJion()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = j
	c.ServeJSON()
}

//@Title 获得活动
//@Description 获得学校发布的信息
//@Success 200 {object} models.activity.Activity
//@Param who query string true 获得哪里的信息，root为学校，grade为年级，class为班级
//@Failure 500 数据库错误
//@Failure 400 输入错误
//@router /activity [get]
func (c *StudentController) GetActivity() {
	name := c.GetString("who")
	var err error
	sess := c.GetSession("username")
	if name == "class" {
		stu := models.Student{Id: sess.(string)}
		name, err = stu.CheckClassTeacher()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
	} else if name == "grade" {
		stu := models.Student{Id: sess.(string)}
		name,err = stu.CheckGradeTeacher()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
	}else if name != "root" {
		beego.Error("输入错误")
		c.Abort("401")
	}
	data, err := models.ShowActivities(name)
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = data
	c.ServeJSON()
}

//@Title 报名活动
//@Description 报名活动，并输入活动的信息,报名后等待老师审核
//@Success 200 {string} "您已经报过名了！" 或 "报名成功！"
//@Param id query string true 活动ID
//@Param message query object false 在活动有额外信息的时候是必填的
//@Failure 500 数据库错误
//@router /jion [get]
func (c *StudentController) SetJion() {
	sess := c.GetSession("username")
	jion := models.Jion{
		ActivityId: c.GetString("id"),
		StudentId:  sess.(string),
		Message:    c.GetString("message"),
	}
	if jion.Message == "" {
		jion.Message = "[]"
	}
	b := jion.Check()
	if b {
		c.Data["json"] = "您已经报过名了！"
	} else {
		jion.Status = "审核中"
		jion.Date = time.Now()
		err := jion.Insert()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
		c.Data["json"] = "报名成功！"
	}
	c.ServeJSON()
}
