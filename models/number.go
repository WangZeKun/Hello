package models


//学生信息
type Number struct {
	Id     string `xorm:"pk char(8)"`              //教育ID
	Name   string `xorm:"varchar(20) notnull"` //学生姓名
	Gender string `xorm:"char(1) notnull"`         //学生性别
	Grade  string `xorm:"char(4) notnull"`         //学生年级
	Class  string `xorm:"char(4) notnull"`         //学生班级
	Avatar string `xorm:"LONGTEXT notnull"`
	Type   string `xorm:"char(7) default 'student' notnull"`
}

//读取学生信息
func (s *Number) Read() (err error) {
	_, err = engine.Get(s)
	return
}

//得到这个学生都参加了什么活动
func (s *Number) ShowWhatJoin() (data []OutStudentJoin, err error) {
	err = engine.SQL("select j.id,a.name,j.date,j.status,a.who_build,a.date as adate from `join` j,activity a where j.activity_id = a.id and j.student_id = ?", s.Id).
		Find(&data)
	return
}
func (s *Number) ShowNotice()(data []Notice,err error)  {
	err = engine.Where("number_id = ?",s.Id).Find(&data)
	return
}

//查找班级管理老师
func (s Number) CheckClassTeacher() (out string, err error) {
	s.Read()
	_,err = engine.Table("number").Where("grade = ? and class = ? and type = 'teacher'",s.Grade,s.Class).
		Cols("id").Get(&out)
	return
}

//查找年级管理老师
func (s Number) CheckGradeTeacher() (out string, err error) {
	s.Read()
	_,err = engine.Table("number").Where("grade = ? and class = ? and type = 'teacher'",s.Grade,"").
		Cols("id").Get(&out)
	return
}

//修改头像
func (s *Number) ChangeAvatar(img string) (err error) {
	_,err = engine.Id(s.Id).Update(s)
	return
}
