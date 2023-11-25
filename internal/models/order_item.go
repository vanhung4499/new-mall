package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint `gorm:"not null"`
	ProductID uint
	Quantity  int
	Price     float64
}

func (OrderItem) TableName() string {
	return "order_items"
}

type OrderItemCreate struct {
	OrderID   uint    `form:"order_id" binding:"required"`
	ProductID uint    `form:"product_id" binding:"required"`
	Quantity  int     `form:"quantity" binding:"required"`
	Price     float64 `form:"price" binding:"required"`
}

func (OrderItemCreate) TableName() string {
	return OrderItem{}.TableName()
}

type OrderItemUpdate struct {
	OrderID   uint    `form:"order_id" binding:"required"`
	ProductID uint    `form:"product_id" binding:"required"`
	Quantity  int     `form:"quantity" binding:"required"`
	Price     float64 `form:"price" binding:"required"`
}

func (OrderItemUpdate) TableName() string {
	return OrderItem{}.TableName()
}
