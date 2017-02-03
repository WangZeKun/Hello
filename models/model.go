package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/gqmms?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Login), new(Student), new(Activity), new(Jion), new(Exam), new(Score))

	// create table
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

//返回哪些活动正在报名
func ShowActivities() ([]Activity, error) {
	o := orm.NewOrm()
	var data []Activity
	_, err := o.QueryTable("activity").Filter("isrecruit", true).All(&data)
	return data, err
}
