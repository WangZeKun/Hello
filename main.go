package main

import (
	_ "hello/routers"

	"github.com/astaxie/beego"
)

func main() {
	// exam := models.Exam{Name:"aaaa"}
	// exam.NewExam()
	// exam.WriteExcel("111")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	//beego.SetStaticPath("/css", "static/css")
	beego.Run()
	/*
		//var can []*models.Canjia
		o := orm.NewOrm()
		var users []User
		o.Raw("SELECT username, password FROM login ").QueryRows(&users)
		for _,i := range users{

			if len(i.Username)==7{
				fmt.Printf("%d",len(i.Username))
				o.Raw("UPDATE login SET `username` = ?,`password`=? WHERE `username` = ?","0"+i.Username,"0"+i.Password,i.Username).Exec()
			}
		}
	*/
}
