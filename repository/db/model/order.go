package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
	BossID    uint   `gorm:"not null"`
	AddressID uint   `gorm:"not null"`
	Num       int    // quantity
	OrderNum  uint64 // order number
	Type      uint   // 1 Unpaid 2 Paid
	Money     float64
}
