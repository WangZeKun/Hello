package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//学生个人活动信息
type Jion struct {
	Id         int       `orm:"auto"`                //Id号
	Status     string    `orm:"size(5)"`             //报名状态
	Date       time.Time `orm:"auto_now;type(date)"` //报名时间
	ActivityId string                                //报名的活动
	StudentId  string                                //学生教育Id
	Message    string                                //额外信息
}

//得到时间
//输出@string:时间
func (c Jion) GetTime() (out string) {
	out = c.Date.Format("2006-01-02")
	return
}

//检查此学生是否已经报名
func (c *Jion) Check() (b bool) {
	o := orm.NewOrm()
	b = o.QueryTable("jion").Filter("activity_id", c.ActivityId).
		Filter("student_id", c.StudentId).Exist()
	return b
}

//报名
func (c *Jion) Insert() (err error) {
	o := orm.NewOrm()
	c.Date = time.Now()
	_, err = o.Insert(c)
	return err
}

//更新状态
func (c *Jion) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(c, "status")
	return err
}

func JionUpdate(js []Jion) error{
	o:=orm.NewOrm()
	_,err := o.Update(&js,"status")
	return err
}
