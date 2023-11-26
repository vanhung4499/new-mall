package models

import "gorm.io/gorm"

const FavoriteEntityName = "Favorite"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserID"`
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"`
}

func (Favorite) TableName() string {
	return "favorites"
}

type FavoriteCreate struct {
	gorm.Model
	UserID    uint `form:"user_id,omitempty"`
	ProductID uint `form:"product_id"`
}

func (FavoriteCreate) TableName() string {
	return Favorite{}.TableName()
}
