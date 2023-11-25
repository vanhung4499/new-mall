package models

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Text string `gorm:"types:text"`
}

func (Notice) TableName() string {
	return "notices"
}
