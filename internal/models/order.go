package models

import "gorm.io/gorm"

const OrderEntityName = "Order"

type Order struct {
	gorm.Model
	UserID      uint    `gorm:"not null"`
	AddressID   uint    `gorm:"not null"`
	Address     Address `gorm:"ForeignKey:AddressID"`
	TotalAmount float64
	OrderItems  []OrderItem
	Status      uint // 1 Unpaid 2 Paid 3 Shipped 4 Completed 5 Cancelled
}

func (Order) TableName() string {
	return "orders"
}

type OrderCreate struct {
	gorm.Model
	UserID      uint    `form:"user_id" binding:"required"`
	AddressID   uint    `form:"address_id" binding:"required"`
	Status      uint    `form:"type" gorm:"default:1" binding:"required"` // 1 Unpaid 2 Paid 3 Shipped 4 Completed 5 Cancelled
	TotalAmount float64 `form:"total_amount" binding:"required"`
	OrderItems  []OrderItemCreate
}

func (OrderCreate) TableName() string {
	return Order{}.TableName()
}
