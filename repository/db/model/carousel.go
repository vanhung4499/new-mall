package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	ImagePath string
	ProductID uint `gorm:"not null"`
}
