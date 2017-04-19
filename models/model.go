package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/gqmms?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Login), new(Student), new(Activity), new(Jion), new(Exam), new(Teacher))

	// create table
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

//得到哪些活动正在报名
func ShowActivities(who string) (data []Activity, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("activity").Filter("isrecruit", true).Filter("who_build", who).All(&data)
	return
}
//得到已经结束的活动
func ShowAllActivities() (data []Activity, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("activity").Filter("isrecruit", false).All(&data)
	return
}

