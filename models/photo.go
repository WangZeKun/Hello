package models

import "github.com/astaxie/beego/orm"

type Photo struct {
	Id int `orm:"auto"`
	ActivityId int
	Photo string
}

func (p *Photo) Insert() (err error) {
	o:=orm.NewOrm()
	_,err = o.Insert(p)
	return
}

func (p *Photo) Delete() (err error){
	o := orm.NewOrm()
	_,err = o.Delete(p)
	return
}