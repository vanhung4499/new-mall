package models

import (
	"errors"
	"gorm.io/gorm"
	"new-mall/pkg/common"
)

const CartEntityName = "Cart"

type Cart struct {
	gorm.Model
	UserID      uint    `gorm:"not null"`
	TotalAmount float64 `gorm:"default:0"`
	CartItems   []CartItem
}

func (Cart) TableName() string {
	return "carts"
}

type CartCreate struct {
	gorm.Model
	UserID uint `form:"user_id" json:"user_id"`
}

func (CartCreate) TableName() string {
	return Cart{}.TableName()
}

var (
	ErrCartIsEmpty = common.NewCustomError(
		errors.New("cart is empty"),
		"cart is empty",
		"ErrCartIsEmpty",
	)
)
