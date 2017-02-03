package models

import "github.com/astaxie/beego/orm"
import "encoding/json"

type Activity struct {
	Id           string `orm:"pk"`
	Name         string
	Number       int
	Introduction string
	Isrecruit    bool
	Message      string
}

func (a *Activity) ShowWhoJion() ([]Jion, error) {
	o := orm.NewOrm()
	var data []Jion
	_, err := o.QueryTable("jion").Filter("activity_id", a.Id).All(&data)
	return data, err
}

func (a *Activity) Read() error {
	o := orm.NewOrm()
	err := o.Read(a)
	return err
}

func (a *Activity) GetMessage() (out []string) {
	json.Unmarshal([]byte(a.Message), &out)
	return
}

func (a *Activity) SetMessage(in []string) (err error) {
	data, err := json.Marshal(in)
	a.Message = string(data)
	return
}
