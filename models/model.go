package models

import (
	"encoding/json"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Name string
	Mess string
	Type string
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/gqmms?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Login), new(Student), new(Activity), new(Jion), new(Exam), new(Teacher))

	// create table
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

//返回哪些活动正在报名
func ShowActivities(who string) (data []Activity, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("activity").Filter("isrecruit", true).Filter("who_build", who).All(&data)
	return
}
func ShowAllActivities() (data []Activity, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("activity").Filter("isrecruit", false).All(&data)
	return
}

func GetJson(in string) (out []Message, err error) {
	if in == "" {
		return
	}
	var data []string
	err = json.Unmarshal([]byte(in), &data)
	if err != nil {
		return
	}
	var m Message
	for _, da := range data {
		err = json.Unmarshal([]byte(da), &m)
		if err != nil {
			return
		}
		out = append(out, m)
	}
	return

}

func SetJson(in []Message) (out string, err error) {
	var data []byte
	var da []string
	for m := range in {
		data, err = json.Marshal(m)
		if err != nil {
			return
		}
		da = append(da, string(data))
	}
	data, err = json.Marshal(da)
	out = string(data)
	return
}
