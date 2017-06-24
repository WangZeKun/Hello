package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	// set default database
	engine, _ = xorm.NewEngine("mysql", "gqmms:pf6zbbF2tt@tcp(127.0.0.1:3306)/gqmms?charset=utf8")
	engine.Sync2(new(Login),new(Activity),new(Number),new(Notice),new(Join))
}

//得到哪些活动正在报名
func ShowActivities(who string, is bool) (data []Activity, err error) {
	err = engine.Where("is_recruit = ?", is).
		And("who_build = ?", who).
		OrderBy("end_start_date").
		Find(&data)
	return
}

//得到这个年纪的班级
func CheckClass(grade string) (class []string, err error) {
	err = engine.Table("number").
		Where("grade = ?", grade).
		Distinct().Cols("class").
		Asc("class").
		Find(&class)
	return
}

//得到这个班的同学
func CheckStudent(grade, class string) (student []Number, err error) {
	err = engine.Where("grade = ?", grade).And("class = ?", class).Find(&student)
	return
}

//获得学分
func getScore(s *Number) (out OutScore, err error) {
	out.Score, err = engine.Where("student_id = ?", s.Id).SumInt(new(Join), "score")
	if err != nil {
		return
	}
	out.Num, err = engine.Where("student_id = ?", s.Id).And("status='审核通过'").Count(new(Join))
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
	_, err = engine.SQL("select sum(score) from `join` where student_id "+
		"in (select id from gqmms.number where class = ?  and grade = ? and type = student ))", class, grade).Get(&out.Score)
	if err != nil {
		return
	}
	_, err = engine.SQL("SELECT count(id) FROM gqmms.join where student_id in "+
		"(select id from gqmms.student where class = ? and grade = ? and type = student ) and status = '审核通过'", class, grade).Get(&out.Num)
	if err != nil {
		return
	}
	out.Class = class
	out.Grade = grade
	return
}

func GetScores(class, grade string) (out []OutScore, err error) {
	if class == "" {
		err = engine.Table("number").Distinct().Cols("class").
			Where("grade = ?", grade).
			OrderBy("class").Asc().
			Iterate(new(string), func(idx int, bean interface{}) error {
			class = bean.(string)
			s, err := getClassScores(class, grade)
			if err != nil {
				return err
			}
			out = append(out, s)
			return nil
		})
	} else {
		err = engine.Table("number").Distinct().
			Where("grade = ?", grade).
			And("class = ?", class).
			Asc("id").
			Iterate(new(Number), func(idx int, bean interface{}) error {
			student := bean.(*Number)
			s, err := getScore(student)
			if err != nil {
				return err
			}
			out = append(out, s)
			return nil
		})
	}
	return
}

type OutScore struct {
	Name  string
	Grade string
	Class string
	Score int64
	Num   int64
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
