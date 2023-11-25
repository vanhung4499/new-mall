package models

import (
	"errors"
	"gorm.io/gorm"
	"new-mall/pkg/common"
)

const CartEntityName = "Cart"

type Cart struct {
	gorm.Model
	UserID    uint
	User      User
	CartItems []CartItem
}

func (Cart) TableName() string {
	return "carts"
}

type CartCreate struct {
	UserID uint `form:"user_id" json:"user_id"`
}

var (
	ErrCartIsEmpty = common.NewCustomError(
		errors.New("cart is empty"),
		"cart is empty",
		"ErrCartIsEmpty",
	)
)
