package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/orm"
)

var mouth_ = map[string]int{"January": 1, "February": 2, "March": 3, "April": 4, "May": 5, "June": 6, "July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12}

//对活动信息进行操作
type Jion struct {
	Id         int       `orm:"auto"`
	Status     string    `orm:"size(1)"`
	Date       time.Time `orm:"auto_now_add;type(date)"`
	ActivityId string
	StudentId  string
	Message    string
}

//得到时间
func (c Jion) GetTime() (out string) {
	out = c.Date.Format("2006-01-02")
	return
}

//检查此学生是否已经报名
func (c *Jion) Check() bool {
	o := orm.NewOrm()
	exist := o.QueryTable("jion").Filter("activity_id", c.ActivityId).Filter("student_id", c.StudentId).Exist()
	return exist
}

//报名
func (c *Jion) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	return err
}

//返回该活动名称
func (c *Jion) CheckActivity() (out string) {
	o := orm.NewOrm()
	o.Raw("SELECT name FROM activity WHERE id = ?", c.ActivityId).QueryRow(&out)
	return out
}

//得到活动的额外信息
//使用json格式写入数据库
func (c *Jion) GetMessage() (out map[string]string) {
	json.Unmarshal([]byte(c.Message), &out)
	return
}

//把活动的额外信息转换为json格式
func (c *Jion) SetMessage(in map[string]string) (err error) {
	data, err := json.Marshal(in)
	c.Message = string(data)
	return
}
