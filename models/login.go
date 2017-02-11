package models

import (
	"github.com/astaxie/beego/orm"
)
//用来对用户的用户名和密码进行操作
type Login struct {
	Username string `orm:"pk"`
	Password string `orm:"size(10)"`
	Who string
}

//在登陆时检查密码是否正确
func (l *Login)Check(password string) bool{
	l.Read()
	if l.Password == password{
		return true
	}else{
		return false
	}
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
