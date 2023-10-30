package model

import "new-mall/global"

type CartItem struct {
	global.Model
	UserId       int `json:"userId" form:"userId" gorm:"column:user_id;type:bigint"`
	ProductId    int `json:"productId" form:"productId" gorm:"column:product_id;type:bigint"`
	ProductCount int `json:"productCount" form:"productCount" gorm:"column:product_count;type:int"`
}

// TableName CartItem
func (CartItem) TableName() string {
	return "cart_items"
}
