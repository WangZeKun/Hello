//用于控制登录的模块
package models

import (
	"github.com/astaxie/beego/orm"
)

//用来对用户的用户名和密码进行操作
type Login struct {
	Username string `orm:"pk"`       //用户名（教育ID)
	Password string `orm:"size(10)"` //密码
	Who      string                  //用户类型
}

//检查密码是否正确
//输入 @string:  密码
//输出 @boolean: 是否正确
func (l *Login) Check(password string) (b bool, err error) {
	err = l.Read()
	if err != nil {
		return false, err
	}
	if l.Password == password {
		b = true
	} else {
		b = false
	}
	return
}

//更新密码
func (l *Login) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(l)
	return err
}

//读取数据
func (l *Login) Read() (err error) {
	o := orm.NewOrm()
	err = o.Read(l)
	return err
}
