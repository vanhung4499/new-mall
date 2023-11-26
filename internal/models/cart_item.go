package models

import "gorm.io/gorm"

const CartItemEntityName = "CartItem"

type CartItem struct {
	gorm.Model
	CartID    uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
}

func (CartItem) TableName() string {
	return "cart_items"
}

type CartItemCreate struct {
	gorm.Model
	CartID    uint    `form:"cart_id" json:"cart_id"`
	ProductID uint    `form:"product_id" json:"product_id"`
	Quantity  int     `form:"quantity" json:"quantity"`
	Price     float64 `form:"price" json:"price"`
}

func (CartItemCreate) TableName() string {
	return CartItem{}.TableName()
}

type CartItemUpdate struct {
	ProductID uint    `form:"product_id" json:"product_id"`
	Quantity  int     `form:"quantity" json:"quantity"`
	Price     float64 `form:"price" json:"price"`
}

func (CartItemUpdate) TableName() string {
	return CartItem{}.TableName()
}
