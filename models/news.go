package models

import (
	"github.com/astaxie/beego/orm"
)

//一个活动的具体信息
type Activity struct {
	Id           int `orm:"auto"`
	Name         string
	Introduction string
	Isrecruit    bool
	Message      string
	Impression   string
	ImagePath    string
	Score        int
	Date         string
	WhoBuild     string `orm:"size(10)"`
}

//返回一个slice 这个活动都谁参加
//输出：这个活动都谁参加
func (a *Activity) ShowWhoJion() (data []OutTeacherJion, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select j.id,s.class,s.grade,s.name,j.status,j.message from jion j,student s where j.student_id = s.id and j.activity_id = ?", a.Id).QueryRows(&data)
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
	a.Isrecruit = true
	if a.ImagePath == "" {
		a.ImagePath = "[]"
	}
	_, err = o.Insert(a)
	return
}

//结束活动
func (a *Activity) EndActivity() (err error) {
	o := orm.NewOrm()
	a.Isrecruit = false
	if a.ImagePath == "" {
		a.ImagePath = "[]"
	}
	_, err = o.Update(a)
	return
}

//删除活动
func (a *Activity) Delete() (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(a)
	if err != nil {
		return
	}
	_, err = o.QueryTable("jion").Filter("activity_id", a.Id).Delete()
	return
}

func (a *Activity) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("jion").Filter("activity_id", a.Id).Update(orm.Params{
		"ischanged": true,
	})
	if err != nil{
		return
	}
	_,err = o.Update(a)
	return
}
