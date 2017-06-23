package models

import "github.com/astaxie/beego/orm"

type Notice struct {
	Id int `orm:"auto"`
	NumberId int
	Message string
}

func (n *Notice) Insert() (err error) {
	o:=orm.NewOrm()
	_,err = o.Insert(n)
	return
}

func (n *Notice) Delete() (err error){
	o := orm.NewOrm()
	_,err = o.Delete(n)
	return
}
