package models

import (
	"time"
)

type Notice struct {
	Id        int `xorm:"autoincr"`
	NumberId  string `xorm:"char(8) notnull"`
	Message   string `xorm:"notnull"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (n *Notice) Insert() (err error) {
	_, err = engine.InsertOne(n)
	return
}

func (n *Notice) Delete() (err error) {
	_, err = engine.Delete(n)
	return
}