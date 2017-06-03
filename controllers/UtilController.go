package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
)

type UtilController struct {
	beego.Controller
}

func (c *UtilController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
}

//跳转
func (c *UtilController) JumpTo() {
	sess := c.GetSession("username")
	se := c.GetSession("select")
	if sess == nil || se == nil {
		c.Redirect("/login", 302)
	} else if se.(string) == "teacher" {
		c.Redirect("/teacher", 302)
	} else if se.(string) == "student" {
		c.Redirect("/student", 302)
	}
}

//检测是否在登录状态
func (c *UtilController) Text() {
	sess := c.GetSession("select")
	if sess == nil {
		c.Data["json"] = sendMessage("没有登录！", nil)
	} else if sess.(string) == "teacher" {
		c.Data["json"] = sendMessage("老师登录！", nil)
	} else if sess.(string) == "student" {
		c.Data["json"] = sendMessage("学生登录！", nil)
	}
	c.ServeJSON()
}

//修改密码
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
		user.Password = c.GetString("newpassword")
		err = user.Update()
		if err != nil {
			beego.Error(err)
			c.Abort("500")
		}
		c.Data["json"] = sendMessage("修改成功", nil)
	}
	c.ServeJSON()
}

//退出登录
func (c *UtilController) Exit() {
	c.DelSession("username")
	c.DelSession("select")
	c.Redirect("/login", 302)
}

//学生登录
func (c *UtilController) Login() {
	user := models.Login{Username: c.GetString("username")}
	b, err := user.Check(c.GetString("password"))
	if err != nil {
		c.Abort("500")
		beego.Informational(err)
	}
	if !b {
		c.Data["json"] = sendMessage("密码错误！", nil)
		c.ServeJSON()
	} else {
		if c.GetString("select") == "学生登陆" && user.Who == "student" {
			sess := c.StartSession()
			sess.Set("username",user.Username)
			sess.Set("select",user.Who)
			beego.Informational(sess.SessionID())
			defer sess.SessionRelease(c.Ctx.ResponseWriter)
			c.Data["json"] = sendMessage("成功，学生登录！",
				map[string]interface{}{
					"id":sess.SessionID(),
					"user":getStudentMessage(user.Username),
				})
		} else if c.GetString("select") == "教师登陆" && user.Who == "teacher" {
			c.SetSession("username", user.Username)
			c.SetSession("select", user.Who)
			sess := c.StartSession()
			beego.Informational(sess.SessionID())
			defer sess.SessionRelease(c.Ctx.ResponseWriter)
			c.Data["json"] = sendMessage("成功，教师登录！",
				map[string]interface{}{
					"id":sess.SessionID(),
					"user":getTeacherMessage(user.Username),
				})
		} else {
			c.Data["json"] = sendMessage("请选择正确的登录用户！", nil)
		}

		c.ServeJSON()
		//c.Ctx.WriteString(str(err))
	}
}

//获取学生信息
func getStudentMessage(id string)(stu models.Student) {
	stu.Id = id
	stu.Read()
	return stu
}

func getTeacherMessage(id string)(tea models.Teacher){
	tea.Id = id
	tea.Read()
	return tea
}