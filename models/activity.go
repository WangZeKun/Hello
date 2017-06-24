package models

import (
	"time"
	"fmt"
)

//一个活动的具体信息
type Activity struct {
	Id           int `xorm:"autoincr pk"`
	Name         string `xorm:"varchar(50) notnull"`
	Introduction string `xorm:"Text notnull"`
	IsRecruit    bool    `xorm:"bool notnull"`
	Message      []string `xorm:"json notnull"`
	Impression   string `xorm:"Text notnull"`
	Date         string `xorm:"DATE notnull"` //报名时间
	EndStartDate time.Time `xorm:"updated DATE notnull"`
	WhoBuild     string `xorm:"varchar(10) notnull"`
}

//返回一个slice 这个活动都谁参加
//输出：这个活动都谁参加
func (a *Activity) ShowWhoJoin() (data []OutTeacherJoin, err error) {
	err = engine.SQL("select j.id,s.class,s.grade,s.name,j.status,j.message from `join` j,`number` s " +
		"where j.student_id = s.id and j.activity_id = ?",a.Id).
		Find(&data)
	return
}

//读取活动信息
func (a *Activity) Read() (err error) {
	_, err = engine.Get(a)
	return
}

//新发布活动
func (a *Activity) Insert() (err error) {
	a.IsRecruit = true
	a.EndStartDate = time.Now()
	_, err = engine.InsertOne(a)
	return
}

//结束活动
func (a *Activity) EndActivity(score int) (err error) {
	session := engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		return
	}
	a.IsRecruit = false
	a.EndStartDate = time.Now()
	_, err = session.ID(a.Id).Update(a)
	_, err = session.ID(a.Id).Cols("is_recruit").Update(a)
	if err != nil {
		fmt.Println(err)
		session.Rollback()
		return
	}
	err = session.Where("activity_id = ?", a.Id).Iterate(new(Join), func(idx int, bean interface{}) error {
		join := bean.(*Join)
		if join.Status == "审核中" {
			join.Status = "审核不通过"
			notice := Notice{
				NumberId: join.StudentId,
				Message:  "您所报名的" + a.Name + "活动已结束，很遗憾，您并没有通过审核。",
			}
			err := notice.Insert()
			if err != nil {
				return err
			}
		} else if join.Status == "审核通过" {
			join.Score = score
			notice := Notice{
				NumberId: join.StudentId,
				Message:  fmt.Sprint("您所报名的", a.Name, "活动已结束，您获得了", score, "学分"),
			}
			err := notice.Insert()
			if err != nil {
				return err
			}
		}
		_, err := session.ID(join.Id).Update(join)
		return err
	})
	if err != nil {
		session.Rollback()
		return
	}
	return session.Commit()
}

//删除活动
func (a *Activity) Delete() (err error) {
	session := engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		return
	}
	_, err = session.Delete(a)
	if err != nil {
		session.Rollback()
		return
	}
	err = session.Table("join").Where("activity_id=?", a.Id).Iterate(new(Join), func(idx int, bean interface{}) error {
		join := bean.(*Join)
		notice := Notice{
			NumberId: join.StudentId,
			Message:  fmt.Sprint("您所报名的", a.Name, "活动已被发起人删除"),
		}
		err := notice.Insert()
		if err != nil {
			return err
		}
		_, err = session.Delete(join)
		return err
	})
	if err != nil {
		session.Rollback()
		return
	}
	return session.Commit()
}

func (a *Activity) Update() (err error) {
	session := engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		return
	}
	_, err = session.ID(a.Id).Update(a)
	if err != nil {
		session.Rollback()
		return
	}
	// TODO 通知学生修改内容，并重新报名
	return session.Commit()
}
