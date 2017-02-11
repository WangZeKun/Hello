package models

import "github.com/astaxie/beego/orm"
import "encoding/json"

//一个活动的具体信息
type Activity struct {
	Id           string `orm:"pk"`
	Name         string
	Introduction string
	Isrecruit    bool
	Message      string
}

//返回一个slice 这个活动都谁参加
func (a *Activity) ShowWhoJion() ([]Jion, error) {
	o := orm.NewOrm()
	var data []Jion
	_, err := o.QueryTable("jion").Filter("activity_id", a.Id).All(&data)
	return data, err
}

//
func (a *Activity) Read() error {
	o := orm.NewOrm()
	err := o.Read(a)
	return err
}

//得到活动的额外信息
//使用json格式写入数据库
func (a *Activity) GetMessage() (out []string) {
	json.Unmarshal([]byte(a.Message), &out)
	return
}

//把活动的额外信息转换为json格式
func (a *Activity) SetMessage(in []string) (err error) {
	data, err := json.Marshal(in)
	a.Message = string(data)
	return
}
