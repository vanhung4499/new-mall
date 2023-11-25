package models

import "gorm.io/gorm"

const CategoryEntityName = "Category"

type Category struct {
	gorm.Model
	CategoryName string
}

func (Category) TableName() string {
	return "categories"
}

type CategoryCreate struct {
	CategoryName string `form:"category_name" binding:"required"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

type CategoryUpdate struct {
	CategoryName string `form:"category_name" binding:"required"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}
