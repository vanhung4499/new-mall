package models

import "gorm.io/gorm"

const AddressEntityName = "Address"

type Address struct {
	gorm.Model
	UserID    uint `gorm:"not null" json:"user_id"`
	Address   string
	Street    string
	City      string
	State     string
	Country   string
	Phone     string
	ZipCode   string
	IsDefault bool `gorm:"default:false"` // Indicates whether this is the default address for the user
}

func (Address) TableName() string {
	return "addresses"
}

type AddressCreate struct {
	UserID    uint   `gorm:"not null" form:"user_id" binding:"required"`
	Address   string `form:"address" binding:"required"`
	Street    string `form:"street" binding:"required"`
	City      string `form:"city" binding:"required"`
	State     string `form:"state" binding:"required"`
	Country   string `form:"country" binding:"required"`
	Phone     string `form:"phone" binding:"required"`
	ZipCode   string `form:"zip_code" binding:"required"`
	IsDefault bool   // Indicates whether this is the default address for the user
}

func (AddressCreate) TableName() string {
	return Address{}.TableName()
}

type AddressUpdate struct {
	UserID    uint   `gorm:"not null" form:"user_id" binding:"required"`
	Address   string `form:"address" binding:"required"`
	Street    string `form:"street" binding:"required"`
	City      string `form:"city" binding:"required"`
	State     string `form:"state" binding:"required"`
	Country   string `form:"country" binding:"required"`
	Phone     string `form:"phone" binding:"required"`
	ZipCode   string `form:"zip_code" binding:"required"`
	IsDefault bool   // Indicates whether this is the default address for the user
}

func (AddressUpdate) TableName() string {
	return Address{}.TableName()
}
