package models

import (
	"time"
)

//学生个人活动信息
type Join struct {
	Id         int       `xorm:"autoincr pk int(11)"`     //Id号
	Status     string    `xorm:"char(5) notnull"`      //报名状态
	Date       time.Time `xorm:"updated DATE notnull"` //报名时间
	ActivityId int       `xorm:"INT notnull"`          //报名的活动
	StudentId  string    `xorm:"char(8) notnull"`      //学生教育Id
	Message    map[string]string `xorm:"json TEXT notnull"` //额外信息
	Score      int       `xorm:"TINYINT(2) default 0 notnull"`          //学分
}

//得到时间
//输出@string:时间
func (c Join) GetTime() (out string) {
	out = c.Date.Format("2006-01-02")
	return
}

//检查此学生是否已经报名
func (c *Join) Check() (b bool, err error) {
	b, err = engine.Where("activity_id = ?", c.ActivityId).Get(c)
	return
}

//报名
func (c *Join) Insert() (err error) {
	_, err = engine.InsertOne(c)
	return err
}

//更新状态
func (c *Join) Update() error {
	_, err := engine.ID(c.Id).Cols("status").Update(c)
	return err
}
