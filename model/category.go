package model

import (
	"new-mall/global"
)

type Category struct {
	global.Model
	CategoryLevel int    `json:"categoryLevel" gorm:""`
	ParentId      int    `json:"parentId" gorm:""`
	CategoryName  string `json:"categoryName" gorm:""`
	CategoryRank  int    `json:"categoryRank" gorm:""`
}

func (Category) TableName() string {
	return "categories"
}
