package models

import "gorm.io/gorm"

const CarouselEntityName = "Carousel"

type Carousel struct {
	gorm.Model
	Title     string
	ImageURL  string
	TargetURL string
	ProductID uint `gorm:"not null"`
}

func (Carousel) TableName() string {
	return "carousels"
}

type CarouselCreate struct {
	gorm.Model
	Title     string `form:"title" binding:"required"`
	ImageURL  string `form:"image_url" binding:"required"`
	TargetURL string `form:"target_url" binding:"required"`
	ProductID uint   `form:"product_id" binding:"required"`
}

func (CarouselCreate) TableName() string {
	return Carousel{}.TableName()
}

type CarouselUpdate struct {
	gorm.Model
	Title     string `form:"title" binding:"required"`
	ImageURL  string `form:"image_url" binding:"required"`
	TargetURL string `form:"target_url" binding:"required"`
	ProductID uint   `form:"product_id" binding:"required"`
}

func (CarouselUpdate) TableName() string {
	return Carousel{}.TableName()
}
