package model

import (
	"new-mall/global"
)

type OrderItem struct {
	global.Model
	OrderId         int    `json:"orderId" form:"orderId" gorm:"column:order_id;type:bigint"`
	ProductId       int    `json:"productId" form:"productId" gorm:"column:product_id;type:bigint"`
	ProductName     string `json:"productName" form:"productName" gorm:"column:product_name;type:varchar(200);"`
	ProductCoverImg string `json:"productCoverImg" form:"productCoverImg" gorm:"column:product_cover_img;type:varchar(200);"`
	SellingPrice    int    `json:"sellingPrice" form:"sellingPrice" gorm:"column:selling_price;type:int"`
	ProductCount    int    `json:"productCount" form:"productCount" gorm:"column:product_count;type:bigint"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
