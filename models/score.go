package models

import "github.com/astaxie/beego/orm"

type Exam struct {
	Id   int `orm:"auto"`
	Name string
}

type Score struct {
	Id        int `orm:"auto"`
	StudentId int
	ExamId    int
	object    string
	score     int
}

func (e *Exam) GetScore(s Student) (data []Score, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("score").Filter("exam_id", e.Id).Filter("student_id", s.Id).All(&data)
	return
}

func (e *Exam) updateScore(s Student){
	o := orm.NewOrm()
	data,_ := e.GetScore(s)
	var exsit = false
	score := Score{StudentId:s.Id, ExamId: e.Id, object: "总分"}
	for _,s1 := range data {
		if s1.object != "总分"{
			score.score += s1.score
		}else{
			exsit=true
		}
	}
	if exsit {
		o.Update(score)
	}else{
		o.Insert(score)
	}
}


