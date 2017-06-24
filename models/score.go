package models
// TODO 老师同意增加学习成绩查询系统，基本没戏
//
//import "github.com/astaxie/beego/orm"
//
//
//
//type Exam struct {
//	Id   int `orm:"auto"`
//	Name string
//}
//
//type Score struct {
//	Id        int `orm:"auto"`
//	StudentId string
//	ExamId    string
//	Object    string
//	Score     string
//}
//
//func (e *Exam) GetScore(s Number) (data []Score, err error) {
//	o := orm.NewOrm()
//	_, err = o.QueryTable("score").Filter("exam_id", e.Id).Filter("student_id", s.Id).All(&data)
//	return
//}
//
//func (e *Exam) updateScore(s Number) (err error) {
//	o := orm.NewOrm()
//	data, err := e.GetScore(s)
//	if err != nil {
//		return
//	}
//	var exsit = false
//	score := Score{StudentId: s.Id, ExamId: string(e.Id), Object: "总分"}
//	for _, s1 := range data {
//		if s1.Object != "总分" {
//			score.Score += s1.Score
//		} else {
//			exsit = true
//		}
//	}
//	if exsit {
//		_, err = o.Update(score)
//	} else {
//		_, err = o.Insert(score)
//	}
//	return
//}
//
//func (e *Exam) NewExam() (err error) {
//	o := orm.NewOrm()
//	_, err = o.Insert(e)
//	o.Read(e)
//	return
//}
//
//func (s *Score) Update() (err error) {
//	o := orm.NewOrm()
//	_, err = o.Update(s)
//	return
//}
//
//func (s *Score) Insert() (err error) {
//	o := orm.NewOrm()
//	_, err = o.Insert(s)
//	return
//}
//
//// func (e *Exam) WriteExcel(filename string) {
//// 	xlFile, err := xlsx.OpenFile("D:MyGo/src/hello/static/file/1.xlsx")
//// 	if err != nil {
//// 	}
//// 	o := orm.NewOrm()
//// 	for _, sheet := range xlFile.Sheets {
//// 		for _, row := range sheet.Rows {
//// 			stu := new(Number)
//// 			for index, cell := range row.Cells {
//// 				switch index {
//// 				case 0:
//// 					stu.Class = getString(cell)+"班"
//// 					print(stu.Class)
//// 				case 1:
//// 					stu.Name = getString(cell)
//// 					print(stu.Name)
//// 				case 2, 3, 4, 5, 6, 7:
//// 					print(string(e.Id))
//// 					err := o.QueryTable("student").Filter("class", stu.Class).Filter("name",stu.Name).One(stu)
//// 					if err == nil {
//// 						o.Insert(&Score{StudentId: stu.Id,Object: getString(sheet.Rows[0].Cells[index]),
//// 							Score: getString(cell)})
//// 					} else {
//// 						fmt.Println(err)
//// 					}
//// 				}
//
//// 			}
//
//// 		}
//// 		break
//// 	}
//// }
//
//// func getString(c *xlsx.Cell) (str string) {
//// 	str, _ = c.String()
//// 	return
//// }
