package models

import (
	"github.com/astaxie/beego/orm"
)

//学生信息
type Student struct {
	Id     string `orm:"pk"`      //教育ID
	Name   string                 //学生姓名
	Gender string `orm:"size(1)"` //学生性别
	Grade  string `orm:"size(3)"` //学生年级
	Class  string `orm:"size(3)"` //学生班级
}

//教师信息
type Teacher struct {
	Id    string `orm:"pk"` //教师用户名
	Class string            //教师班级
	Grade string            //教师年级
}

type OutScore struct {
	Name  string
	Grade string
	Class string
	Score string
	Num   string
}

//读取教师信息
func (t *Teacher) Read() (err error) {
	o := orm.NewOrm()
	err = o.Read(t)
	return
}

//读取学生信息
func (s *Student) Read() (err error) {
	if s.Id != "" {
		o := orm.NewOrm()
		err = o.Read(s)
		return
	} else {
		o := orm.NewOrm()
		err = o.QueryTable("student").Filter("class", s.Class).Filter("grade", s.Grade).Filter("Name", s.Name).One(s)
		return
	}
}

//得到这个学生都参加了什么活动
func (s *Student) ShowWhatJion() (data []OutStudentJion, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select j.id,a.name,j.date,j.status,a.who_build from jion j,activity a where j.activity_id = a.id and j.student_id = ?", s.Id).QueryRows(&data)
	return
}


//查找班级管理老师
func (s Student) CheckClassTeacher() (out string, err error) {
	o := orm.NewOrm()
	s.Read()
	o.Raw("select id from teacher where grade = ? and class = ?", s.Grade, s.Class).QueryRow(&out)
	return
}

//查找年级管理老师
func (s Student) CheckGradeTeacher() (out string, err error) {
	o := orm.NewOrm()
	s.Read()
	o.Raw("select id from teacher where grade = ? and class = ?", s.Grade, "").QueryRow(&out)
	return
}

//得到这个年纪的班级
func CheckClass(grade string) (class []string, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select distinct class from student where grade = ? order by class", grade).QueryRows(&class)
	return
}

//得到这个班的同学
func CheckStudent(grade, class string) (student []Student, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select id,name from student where grade = ? and class = ?", grade, class).QueryRows(&student)
	return
}

//获得学分
func getScore(s *Student) (out OutScore, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT sum(score) FROM gqmms.activity where id in (select activity_id from jion where student_id = ?)", s.Id).QueryRow(&out.Score)
	if err != nil {
		return
	}
	err = o.Raw("SELECT count(id) FROM gqmms.jion where student_id = ?", s.Id).QueryRow(&out.Num)
	if err != nil {
		return

	}
	if s.Class == "" {
		err = s.Read()
		if err != nil {
			return
		}
	}
	out.Class = s.Class
	out.Grade = s.Grade
	out.Name = s.Name
	return
}

//得到班级总分
func GetClassScores(class, grade string) (out OutScore, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT sum(score) FROM gqmms.activity where id in (select activity_id from jion where student_id in (select id from gqmms.student where class = ?  and grade = ? ))", class, grade).QueryRow(&out.Score)
	if err != nil {
		return
	}
	err = o.Raw("SELECT count(id) FROM gqmms.jion where student_id in (select id from gqmms.student where class = ? and grade = ? )", class, grade).QueryRow(&out.Num)
	if err != nil {
		return
	}
	out.Class = class
	out.Grade = grade
	return
}

func GetScores(class, grade string) (o []OutScore, err error) {
	or := orm.NewOrm()
	if class == "" {
		var classes []string
		_, err = or.Raw("select distinct class from gqmms.student where grade =? order by class", grade).QueryRows(&classes)
		if err != nil {
			return
		}
		for _, i := range classes {
			var s OutScore
			s, err = GetClassScores(i, grade)
			if err != nil {
				return
			}
			o = append(o, s)
		}
	} else {
		var students []Student
		_, err = or.Raw("select * from gqmms.student where class = ? and grade = ? order by id ", class, grade).QueryRows(&students)
		if err != nil {
			return
		}
		for _, i := range students {
			var s OutScore
			s, err = getScore(&i)
			if err != nil {
				return
			}
			o = append(o, s)

		}
	}
	return
}
