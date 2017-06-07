package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//func init() {
//	// set default database
//	orm.RegisterDataBase("default", "mysql", "gqmms:pf6zbbF2tt@tcp(127.0.0.1:3306)/gqmms?charset=utf8", 30)
//
//	// register model
//	orm.RegisterModel(new(Login), new(Student), new(Activity), new(Jion), new(Exam), new(Teacher))
//
//	// create table
//	orm.RunSyncdb("default", false, true)
//	orm.Debug = true
//}

//得到哪些活动正在报名
func ShowActivities(who string) (data []Activity, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("activity").Filter("isrecruit", true).Filter("who_build", who).All(&data)
	return
}

//得到已经结束的活动
func ShowAllActivities() (data []Activity, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("activity").Filter("isrecruit", false).All(&data)
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

type OutScore struct {
	Name  string
	Grade string
	Class string
	Score string
	Num   string
}


type OutTeacherJion struct {
	Id      string
	Class   string
	Grade   string
	Status  string
	Name    string
	Message string
}

type OutStudentJion struct {
	Id       string
	Date     string
	Name     string
	Status   string
	WhoBuild string
	Adate    string
}
