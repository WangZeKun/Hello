package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "gqmms:pf6zbbF2tt@tcp(47.94.91.118:3306)/gqmms?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Login),
		new(Student),
		new(Activity),
		new(Join),
		new(Exam),
		new(Teacher),
		new(Photo),
	)

	// create table
	orm.Debug = true
}

//得到哪些活动正在报名
func ShowActivities(who string, is bool) (data []Activity, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id", "name").
		From("activity").
		Where("isrecruit=?").
		And("who_build=?").
		OrderBy("endstartDate").Desc()
	_, err = o.Raw(qb.String(), is, who).QueryRows(&data)
	return
}

//得到这个年纪的班级
func CheckClass(grade string) (class []string, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("distinct class").
		From("student").
		Where("grade =?").
		OrderBy("class").Asc()
	_, err = o.Raw(qb.String(), grade).QueryRows(&class)
	return
}

//得到这个班的同学
func CheckStudent(grade, class string) (student []Student, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id",
		"name").
		From("student").
		Where("grade=?").
		And("class = ?").
		OrderBy("class").Desc()
	_, err = o.Raw("select id,name from student where grade = ? and class = ?", grade, class).QueryRows(&student)
	return
}

//获得学分
func getScore(s *Student) (out OutScore, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT sum(score) FROM gqmms.join where student_id =?", s.Id).
		QueryRow(&out.Score)
	if err != nil {
		return
	}
	err = o.Raw("SELECT count(id) FROM gqmms.join where student_id = ? and status = '审核通过'", s.Id).QueryRow(&out.Num)
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
func getClassScores(class, grade string) (out OutScore, err error) {
	o := orm.NewOrm()
	err = o.Raw("select sum(score) from `join` where student_id "+
		"in (select id from gqmms.student where class = ?  and grade = ? ))", class, grade).QueryRow(&out.Score)
	if err != nil {
		return
	}
	err = o.Raw("SELECT count(id) FROM gqmms.join where student_id in " +
		"(select id from gqmms.student where class = ? and grade = ? ) and status = '审核通过'", class, grade).QueryRow(&out.Num)
	if err != nil {
		return
	}
	out.Class = class
	out.Grade = grade
	return
}

func GetScores(class, grade string) (out []OutScore, err error) {
	or := orm.NewOrm()
	if class == "" {
		var classes []string
		_, err = or.Raw("select distinct class from gqmms.student where grade =? order by class", grade).QueryRows(&classes)
		if err != nil {
			return
		}
		for _, i := range classes {
			var s OutScore
			s, err = getClassScores(i, grade)
			if err != nil {
				return
			}
			out = append(out, s)
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
			out = append(out, s)

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

type OutTeacherJoin struct {
	Id      string
	Class   string
	Grade   string
	Status  string
	Name    string
	Message string
}

type OutStudentJoin struct {
	Id       string
	Date     string
	Name     string
	Status   string
	WhoBuild string
	Adate    string
}
