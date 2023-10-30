package model

import (
	"new-mall/global"
)

type Carousel struct {
	global.Model
	CarouselUrl  string `json:"carouselUrl" form:"carouselUrl" gorm:"column:carousel_url;type:varchar(100);"`
	RedirectUrl  string `json:"redirectUrl" form:"redirectUrl" gorm:"column:redirect_url;type:varchar(100);"`
	CarouselRank int    `json:"carouselRank" form:"carouselRank" gorm:"column:carousel_rank;type:int"`
	CreateUser   int    `json:"createUser" form:"createUser" gorm:"column:create_user;type:int"`
	UpdateUser   int    `json:"updateUser" form:"updateUser" gorm:"column:update_user;type:int"`
}

// TableName Carousel
func (Carousel) TableName() string {
	return "carousels"
}
