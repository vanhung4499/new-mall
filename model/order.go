package model

import (
	"new-mall/global"
	"time"
)

type Order struct {
	global.Model
	OrderNo     string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;type:varchar(20);"`
	UserId      int       `json:"userId" form:"userId" gorm:"column:user_id;type:bigint"`
	TotalPrice  int       `json:"totalPrice" form:"totalPrice" gorm:"column:total_price;type:int"`
	PayStatus   int       `json:"payStatus" form:"payStatus" gorm:"column:pay_status;type:tinyint"`
	PayType     int       `json:"payType" form:"payType" gorm:"column:pay_type;type:tinyint"`
	PayTime     time.Time `json:"payTime" form:"payTime" gorm:"column:pay_time;type:datetime"`
	OrderStatus int       `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;type:tinyint"`
	ExtraInfo   string    `json:"extraInfo" form:"extraInfo" gorm:"column:extra_info;type:varchar(100);"`
}

// TableName Order
func (Order) TableName() string {
	return "orders"
}
