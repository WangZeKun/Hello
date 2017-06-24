package models

type Photo struct {
	Id int `xorm:"autoincr"`
	ActivityId int
	Photo string `xorm:"LONGTEXT"`
}

func (p *Photo) Insert() (err error) {
	_,err = engine.Insert(p)
	return
}

func GetPhotos(activityId int)(photos []Photo,err error)  {
	err = engine.Where("activity_id = ?",activityId).Find(&photos)
	return
}

