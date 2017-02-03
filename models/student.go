package models

import (
	"github.com/astaxie/beego/orm"
)

//用来存储学生信息
type Student struct {
	Id        string `orm:"pk"`
	Name      string
	Gender    string `orm:"size(1)"`
	Section   string `orm:"size(2)"`
	Grade     string `orm:"size(3)"`
	Class     string `orm:"size(3)"`
	Telephone string
	Qq        string
	WeChat    string
	Jianjie   string
}

//读取学生信息
func (s *Student) Read() error {
	o := orm.NewOrm()
	err := o.Read(s)
	return err
}

//更新学生信息
func (s *Student) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(s)
	return err
}

//查看学生参加的活动
func (s *Student) ShowJion() ([]Jion, error) {
	o := orm.NewOrm()
	var data []Jion
	_, err := o.QueryTable("jion").Filter("student_id", s.Id).All(&data)
	return data, err
}
