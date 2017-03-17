package models

import (
	"encoding/json"
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
	WhoBuild     string `orm:"size(10)"`
}

//返回一个slice 这个活动都谁参加
func (a *Activity) ShowWhoJion() (data []OutTeacherJion, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select j.id,s.class,s.grade,s.name,j.status from jion j,student s where j.student_id = s.id and j.activity_id = ?", a.Id).QueryRows(&data)
	return
}

//
func (a *Activity) Read() (err error) {
	o := orm.NewOrm()
	err = o.Read(a)
	return
}

//得到活动的额外信息
//使用json格式从数据库读出
func (a *Activity) GetMessage() (out []Message, err error) {
	if a.Message == "" {
		return
	}
	out, err = GetJson(a.Message)
	return
}

func (a *Activity) GetImagePath() (imagepath []string, err error) {
	message, err := GetJson(a.ImagePath)
	if err != nil {
		return
	}
	for _, m := range message {
		imagepath = append(imagepath, m.Mess)
	}
	return
}

//把活动的额外信息转换为json格式
func (a *Activity) SetMessage(message []Message) (err error) {
	a.Message, err = SetJson(message)
	return
}

func (a *Activity) SetImagePath(in string) (err error) {
	var inn []string
	err = json.Unmarshal([]byte(in), &inn)
	if err != nil {
		return
	}
	var mess []Message
	for _, i := range inn {
		mess = append(mess, Message{Name: "ImagePath", Mess: i})
		a.ImagePath, err = SetJson(mess)
	}
	return

}

func (a *Activity) Insert() (err error) {
	o := orm.NewOrm()
	a.Isrecruit = true
	if a.ImagePath == "" {
		a.ImagePath = "[]"
	}
	_, err = o.Insert(a)
	return
}

func (a *Activity) Update() (err error) {
	o := orm.NewOrm()
	a.Isrecruit = false
	if a.ImagePath == "" {
		a.ImagePath = "[]"
	}
	_, err = o.Update(a)
	return
}

func (a *Activity) Delete() (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("jion").Filter("activity_id", a.Id).Delete()
	if err != nil {
		return
	}
	_, err = o.Delete(a)
	return
}
