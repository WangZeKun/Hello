//用于控制登录的模块
package models

//用来对用户的用户名和密码进行操作
type Login struct {
	Username string `xorm:"pk char(8)"` //用户名（教育ID)
	Password string `xorm:"varchar(20) notnull"`   //密码
	Who      string `xorm:"char(7) notnull"`     //用户类型
}

//检查密码是否正确
//输入 @string:  密码
//输出 @boolean: 是否正确
func (l *Login) Check() (b bool, err error) {
	b, err = engine.Get(l)
	return
}

//更新密码
func (l *Login) Update() (err error) {
	_, err = engine.ID(l.Username).Cols("password").Update(l)
	return err
}
