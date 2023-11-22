package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Name    string `gorm:"types:varchar(20) not null"`
	Phone   string `gorm:"types:varchar(11) not null"`
	Address string `gorm:"types:varchar(50) not null"`
}
