package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//学生个人活动信息
type Join struct {
	Id         int       `orm:"auto"`                //Id号
	Status     string    `orm:"size(5)"`             //报名状态
	Date       time.Time `orm:"auto_now;type(date)"` //报名时间
	ActivityId string                                //报名的活动
	StudentId  string                                //学生教育Id
	Message    string                                //额外信息
	Score      int                                   //学分
}

//得到时间
//输出@string:时间
func (c Join) GetTime() (out string) {
	out = c.Date.Format("2006-01-02")
	return
}

//检查此学生是否已经报名
func (c *Join) Check() (b bool) {
	o := orm.NewOrm()
	b = o.QueryTable("join").Filter("activity_id", c.ActivityId).
		Filter("student_id", c.StudentId).Exist()
	return b
}

//报名
func (c *Join) Insert() (err error) {
	o := orm.NewOrm()
	c.Date = time.Now()
	_, err = o.Insert(c)
	return err
}

//更新状态
func (c *Join) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(c, "status")
	return err
}

func UpdateJoins(js []Join) (err error){
	o := orm.NewOrm()
	//_,err := o.Update(&js,"status","score")
	for _,x := range js{
		_,err = o.Update(&x,"status","score")
		if err!=nil{
			break
		}
	}
	return
}