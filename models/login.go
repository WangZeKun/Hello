package models

import (
	"github.com/astaxie/beego/orm"
)
//用来对用户的用户名和密码进行操作
type Login struct {
	Username string `orm:"pk"`
	Password string `orm:"size(10)"`
	Who string `orm:"-"`
}

//在登陆时检查密码是否正确
func (l *Login)Check() bool{
	o := orm.NewOrm()
	var exist bool
	if l.Who =="学生登陆"{
		exist= o.QueryTable("login_student").Filter("username",l.Username).Filter("password",l.Password).Exist()
	}else{
		exist= o.QueryTable("login_teacher").Filter("username",l.Username).Filter("password",l.Password).Exist()
	}
	return exist
}

//更新密码
func (l *Login)Update() error{
	o := orm.NewOrm()
	_,err := o.Update(l)
	return err
}

//读取数据
func (l *Login) Read() error {
	o := orm.NewOrm()
	err := o.Read(l)
	return err
}


func (l *Login) TableName() (out string) {
    if l.Who == "学生登陆"{
		out = "login_student"
	}else{
		out = "login_teacher"
	}
	return 
}