package controllers

import (
	"hello/models"

	"encoding/json"
	"time"

	"github.com/astaxie/beego"
)

//学生端API
type StudentController struct {
	beego.Controller
}

func (c *StudentController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "https://www.gqmms.wang")
}

//@Title 得到报名活动信息
//@Description 获取学生参加活动的信息
//@Success 200 {object} models.model.OutStudentJoin
//@Failure 500 数据库错误
//@router /canjia [get]
func (c *StudentController) GetCanjia() {
	sess := c.GetSession("username")
	stu := models.Number{Id: sess.(string)}
	j, err := stu.ShowWhatJoin()
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
		stu := models.Number{Id: sess.(string)}
		name, err = stu.CheckClassTeacher()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
	} else if name == "grade" {
		stu := models.Number{Id: sess.(string)}
		name, err = stu.CheckGradeTeacher()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
	} else if name != "root" {
		beego.Error("输入错误")
		c.Abort("401")
	}
	data, err := models.ShowActivities(name, true)
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
//@Param message query Array<string> false 在活动有额外信息的时候是必填的
//@Failure 500 数据库错误
//@router /join [get]
func (c *StudentController) Setjoin() {
	sess := c.GetSession("username")
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	join := models.Join{
		ActivityId: id,
		StudentId:  sess.(string),
	}
	err = json.Unmarshal([]byte(c.GetString("message")), &join.Message)
	b, err := join.Check()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	if b {
		c.Data["json"] = "您已经报过名了！"
	} else {
		join.Status = "审核中"
		join.Date = time.Now()
		err := join.Insert()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
		c.Data["json"] = "报名成功！"
	}
	c.ServeJSON()
}

//@Title 获得消息
//@Description 获得消息
//@Success 200 {string} "成功！"
//@Failure 500 数据库错误
//@router /notice [get]
func (c *StudentController) GetNotices() {
	sess := c.GetSession("username")
	stu := models.Number{Id: sess.(string)}
	data, err := stu.ShowNotice()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = data
	c.ServeJSON()
}

//@Title 已读消息
//@Success 200 {string} "成功！"
//@Param id query string true 消息ID
//@Failure 500 数据库错误参数错误
//@Failure 400
//@router /readNotice [get]
func (c *StudentController) ReadNotices() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
		c.Abort("400")
	}
	n := models.Notice{Id: id}
	err = n.Delete()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = "成功！"
	c.ServeJSON()
}
