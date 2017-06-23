package models

import (
	"github.com/astaxie/beego/orm"
)

//学生信息
type Student struct {
	Id     string `orm:"pk"`      //教育ID
	Name   string                 //学生姓名
	Gender string `orm:"size(1)"` //学生性别
	Grade  string `orm:"size(3)"` //学生年级
	Class  string `orm:"size(3)"` //学生班级
}

//教师信息
type Teacher struct {
	Id    string `orm:"pk"` //教师用户名
	Class string            //教师班级
	Grade string            //教师年级
}



//读取教师信息
func (t *Teacher) Read() (err error) {
	o := orm.NewOrm()
	err = o.Read(t)
	return
}

//读取学生信息
func (s *Student) Read() (err error) {
	if s.Id != "" {
		o := orm.NewOrm()
		err = o.Read(s)
		return
	} else {
		o := orm.NewOrm()
		err = o.QueryTable("student").Filter("class", s.Class).Filter("grade", s.Grade).Filter("Name", s.Name).One(s)
		return
	}
}

//得到这个学生都参加了什么活动
func (s *Student) ShowWhatJoin() (data []OutStudentJoin, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select j.id,a.name,j.date,j.status,a.who_build,a.date as adate from `join` j,activity a where j.activity_id = a.id and j.student_id = ?", s.Id).
		QueryRows(&data)
	return
}


//查找班级管理老师
func (s Student) CheckClassTeacher() (out string, err error) {
	o := orm.NewOrm()
	s.Read()
	o.Raw("select id from teacher where grade = ? and class = ?", s.Grade, s.Class).QueryRow(&out)
	return
}

//查找年级管理老师
func (s Student) CheckGradeTeacher() (out string, err error) {
	o := orm.NewOrm()
	s.Read()
	o.Raw("select id from teacher where grade = ? and class = ?", s.Grade, "").QueryRow(&out)
	return
}

//修改头像
func (s *Student)ChangeAvatar(img string)(err error){
	o := orm.NewOrm()
	_,err = o.Update(s,"avatar")
	return
}
