package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

var mouth_ = map[string]int{"January": 1, "February": 2, "March": 3, "April": 4, "May": 5, "June": 6, "July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12}

type Jion struct {
	Id         int    `orm:"auto"`
	Status     string `orm:"size(1)"`
	Date       time.Time
	ActivityId string
	StudentId  string
	Message    string
}

func (c Jion) GetTime() string {
	year, mouth, day := c.Date.Date()
	_mouth := mouth_[mouth.String()]
	out := fmt.Sprintf("%d-%d-%d", year, _mouth, day)
	return out
}

func (c *Jion) Check() bool {
	o := orm.NewOrm()
	exist := o.QueryTable("jion").Filter("activity_id", c.ActivityId).Filter("student_id", c.StudentId).Exist()
	return exist
}

func (c *Jion) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	return err
}

func (c *Jion) CheckActivity() (out string) {
	o := orm.NewOrm()
	o.Raw("SELECT name FROM activity WHERE id = ?", c.ActivityId).QueryRow(&out)
	return out
}

func (c *Jion) GetMessage() (out map[string]string) {
	json.Unmarshal([]byte(c.Message), &out)
	return
}

func (c *Jion) SetMessage(in map[string]string) (err error) {
	data, err := json.Marshal(in)
	c.Message = string(data)
	return
}
