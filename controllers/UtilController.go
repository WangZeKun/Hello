package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

//工具类 API
type UtilController struct {
	beego.Controller
}

func (c *UtilController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "https://www.gqmms.wang")
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
//@Success 200  {string} Message 原密码错误！或修改成功！
//@Failure 500 数据库错误
//@Param password formData string true "密码"
//@Param newPassword formData string true "新密码"
//@router /change [post]
func (c *UtilController) Change() {
	sess := c.GetSession("username")
	user := models.Login{Username: sess.(string), Password: c.GetString("password")}
	b, err := user.Check()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	if !b {
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
	c.ServeJSON()
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
		beego.Error(err)
		c.Abort("400")
	}
	data := models.Activity{Id: i}
	err = data.Read()
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = data
	c.ServeJSON()
}

//@Title 学生登录
//@Description 用来登录并加上cookie
//@Success 200 {object} controllers.sendMessage.Send
//@Param username formData string true "教育ID"
//@Param password formData string true "密码"
//@Param select formData string true "用户类型"
//@Failure 500 数据库错误
//@router /login [post]
func (c *UtilController) Login() {
	user := models.Login{Username: c.GetString("username"), Password: c.GetString("password")}
	b, err := user.Check()
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
					"user": getNumberMessage(user.Username, "student"),
				})
		} else if c.GetString("select") == "教师登陆" && user.Who == "teacher" {
			c.SetSession("username", user.Username)
			c.SetSession("select", user.Who)
			sess := c.StartSession()
			defer sess.SessionRelease(c.Ctx.ResponseWriter)
			c.Data["json"] = sendMessage("成功，教师登录！",
				map[string]interface{}{
					"id":   sess.SessionID(),
					"user": getNumberMessage(user.Username, "teacher"),
				})
		} else {
			c.Data["json"] = sendMessage("请选择正确的登录用户！", nil)
		}

		c.ServeJSON()
	}
}

//@Title 修改头像
//@Description 修改用户头像
//@Param img formData string true "头像图片"
//@Failure 500 数据库错误
//router /avatar [post]
func (c *UtilController) ChangeAvatar() {
	stu := models.Number{Id: c.GetString("username")}
	err := stu.ChangeAvatar(c.GetString("img"))
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["json"] = "成功！"
	c.ServeJSON()
}


//@Title 获取信息
//@Description
//@Param id Query string true 用户ID
//@Param type Query string true 用户种类
//Failure 500 数据库错误
//router /person [get]
func (c *UtilController) GetPerson(){
	stu := getNumberMessage(c.GetString("id"),c.GetString("type"))
	c.Data["json"] = stu
	c.ServeJSON()
}

//获取学生信息
func getNumberMessage(id, t string) (stu models.Number) {
	stu.Id = id
	stu.Type = t
	stu.Read()
	return stu
}
