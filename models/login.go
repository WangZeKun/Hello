package models

import (
	"github.com/astaxie/beego/orm"
)

type Login struct {
	Username string `orm:"pk"`
	Password string `orm:"size(10)"`
}


func (l *Login)Check() bool{
	o := orm.NewOrm()
	exist:= o.QueryTable("login").Filter("username",l.Username).Filter("password",l.Password).Exist()
	return exist
}

func (l *Login)Update() error{
	o := orm.NewOrm()
	_,err := o.Update(l)
	return err
}

func (l *Login) Read() error {
	o := orm.NewOrm()
	err := o.Read(l)
	return err
}
