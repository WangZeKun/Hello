package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//一个活动的具体信息
type Activity struct {
	Id           int `orm:"auto"`
	Name         string
	Introduction string
	IsRecruit    bool
	Message      string
	Impression   string
	Date         string
	EndStartDate time.Time `orm:"auto_now;type(date)"` //报名时间
	WhoBuild     string `orm:"size(10)"`
	Photos       []*Photo `orm:"reverse(many)"`
}

//返回一个slice 这个活动都谁参加
//输出：这个活动都谁参加
func (a *Activity) ShowWhoJoin() (data []OutTeacherJoin, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select j.id,s.class,s.grade,s.name,j.status,j.message from `join` j,`student` s where j.student_id = s.id and j.activity_id = ?", a.Id).QueryRows(&data)
	return
}

//读取活动信息
func (a *Activity) Read() (err error) {
	o := orm.NewOrm()
	err = o.Read(a)
	return
}

//新发布活动
func (a *Activity) Insert() (err error) {
	o := orm.NewOrm()
	a.IsRecruit = true
	a.EndStartDate = time.Now()
	_, err = o.Insert(a)
	return
}

//结束活动
func (a *Activity) EndActivity(score int) (err error) {
	o := orm.NewOrm()
	a.IsRecruit = false
	a.EndStartDate = time.Now()
	_, err = o.Update(a)
	if err != nil {
		return
	}
	js := []Join{}
	_, err = o.Raw("select * from `join` where activity_id = ?", a.Id).QueryRows(&js)
	for i := range js {
		if js[i].Status == "审核中" {
			js[i].Status = "审核不通过"
		} else if js[i].Status == "审核通过" {
			js[i].Score = score
		}
	}
	err = UpdateJoins(js)
	if err != nil {
		a.IsRecruit = true
		o.Update(a)
	}
	return
}

//删除活动
func (a *Activity) Delete() (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(a)
	if err != nil {
		return
	}
	_, err = o.QueryTable("join").Filter("activity_id", a.Id).Delete()
	return
}

func (a *Activity) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("join").Filter("activity_id", a.Id).Update(orm.Params{
		"ischanged": true,
	})
	if err != nil {
		return
	}
	_, err = o.Update(a)
	return
}
