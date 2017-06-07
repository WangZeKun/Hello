package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
)

//工具类 API
type UtilController struct {
	beego.Controller
}

func (c *UtilController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
}

//@Title 监测登录
//@Description 检测是否在登录状态
//@Success 200  {string}  "没有登录！"or"老师登录！"or"学生登录！"
//@router /text [get]
func (c *UtilController) Text() {
	sess := c.GetSession("select")
	if sess == nil {
		c.Data["json"] = "没有登录！"
	} else if sess.(string) == "teacher" {
		c.Data["json"] = "老师登录！"
	} else if sess.(string) == "student" {
		c.Data["json"] = "学生登录！"
	}
	c.ServeJSON()
}

//@Title 修改密码
//@Description 修改密码。。。。。。
//@Success 200 {string} "原密码错误！"or"修改成功"
//@Failure 500 数据库错误
//@Param password formData string true "密码"
//@Param newPassword formData string true "新密码"
//@router /change [post]
func (c *UtilController) Change() {
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
		user.Password = c.GetString("newPassword")
		err = user.Update()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
		c.Data["json"] = "修改成功"
	}
	c.ServeJSON()
}

//@Title 退出登录
//@Description 退出登录，删除cookie
//@Success 200 {void} 退出登录
//@router /exit [get]
func (c *UtilController) Exit() {
	c.DelSession("username")
	c.DelSession("select")
	c.Redirect("/login", 302)
}

//@Title 获得信息
//@Description 活动单一活动的信息
//@Success 200 {object} models.activity.Activity
//@Param id query int true "活动ID"
//@Failure 400 输入的不是int
//@Failure 500 数据库错误
//@router /single [get]
func (c *UtilController) GetSingle() {
	i, err := c.GetInt("id")
	if err != nil {
		c.Abort("400")
		beego.Error(err)
	}
	data := models.Activity{Id: i}
	err = data.Read()
	if err != nil {
		c.Abort("500")
		beego.Error(err)
	}
	c.Data["json"] = data
	c.ServeJSON()
}

//@Title 学生登录
//@Description 用来登录并加上cookie
//@Success 200 {object} controllers.sendMessage.Send
//@Param username formData string true "教育ID"
//@Param password formData string true "密码"
//@Param selsect formData string true "用户类型"
//@Failure 500 数据库错误
//@router /login [post]
func (c *UtilController) Login() {
	user := models.Login{Username: c.GetString("username")}
	b, err := user.Check(c.GetString("password"))
	if err != nil {
		c.Abort("500")
		beego.Error(err)
	}
	if !b {
		c.Data["json"] = sendMessage("密码错误！", nil)
		c.ServeJSON()
	} else {
		if c.GetString("select") == "学生登陆" && user.Who == "student" {
			sess := c.StartSession()
			sess.Set("username", user.Username)
			sess.Set("select", user.Who)
			defer sess.SessionRelease(c.Ctx.ResponseWriter)
			c.Data["json"] = sendMessage("成功，学生登录！",
				map[string]interface{}{
					"id":   sess.SessionID(),
					"user": getStudentMessage(user.Username),
				})
		} else if c.GetString("select") == "教师登陆" && user.Who == "teacher" {
			c.SetSession("username", user.Username)
			c.SetSession("select", user.Who)
			sess := c.StartSession()
			defer sess.SessionRelease(c.Ctx.ResponseWriter)
			c.Data["json"] = sendMessage("成功，教师登录！",
				map[string]interface{}{
					"id":   sess.SessionID(),
					"user": getTeacherMessage(user.Username),
				})
		} else {
			c.Data["json"] = sendMessage("请选择正确的登录用户！", nil)
		}

		c.ServeJSON()
	}
}

//获取学生信息
func getStudentMessage(id string) (stu models.Student) {
	stu.Id = id
	stu.Read()
	return stu
}

func getTeacherMessage(id string) (tea models.Teacher) {
	tea.Id = id
	tea.Read()
	return tea
}
