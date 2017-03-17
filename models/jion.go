package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//对活动信息进行操作
type Jion struct {
	Id         int       `orm:"auto"`
	Status     string    `orm:"size(5)"`
	Date       time.Time `orm:"auto_now;type(date)"`
	ActivityId string
	StudentId  string
	Message    string
}

type OutTeacherJion struct {
	Id     string
	Class  string
	Grade  string
	Status string
	Name   string
}

type OutStudentJion struct {
	Id     string
	Date   string
	Name   string
	Status string
}

//得到时间
func (c Jion) GetTime() (out string) {
	out = c.Date.Format("2006-01-02")
	return
}

//检查此学生是否已经报名
func (c *Jion) Check() bool {
	o := orm.NewOrm()
	exist := o.QueryTable("jion").Filter("activity_id", c.ActivityId).
		Filter("student_id", c.StudentId).Exist()
	return exist
}

//报名
func (c *Jion) Insert() error {
	o := orm.NewOrm()
	c.Date = time.Now()
	_, err := o.Insert(c)
	return err
}

//更改状态
func (c *Jion) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(c, "status")
	return err
}

//返回该活动名称
func (c *Jion) CheckActivity() (out string, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT name FROM activity WHERE id = ?", c.ActivityId).QueryRow(&out)
	return
}

//返回报名活动的同学
func (c *Jion) CheckStudent() (out Student, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT * FROM student WHERE id = ?", c.StudentId).QueryRow(&out)
	return
}

//得到活动的额外信息
//使用json格式写入数据库
func (c *Jion) GetMessage() (out []Message, err error) {
	if c.Message == "" {
		return
	}
	out, err = GetJson(c.Message)
	return
}

//把活动的额外信息转换为json格式
func (c *Jion) SetMessage(in []Message) (err error) {
	c.Message, err = SetJson(in)
	return
}
