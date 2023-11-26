package models

import "gorm.io/gorm"

const AddressEntityName = "Address"

type Address struct {
	gorm.Model
	UserID    uint
	No        string
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
	gorm.Model
	UserID    uint   `form:"-"`
	No        string `form:"no"`
	Street    string `form:"street"`
	City      string `form:"city"`
	State     string `form:"state"`
	Country   string `form:"country"`
	Phone     string `form:"phone"`
	ZipCode   string `form:"zip_code"`
	IsDefault bool   `form:"is_default"` // Indicates whether this is the default address for the user
}

func (AddressCreate) TableName() string {
	return Address{}.TableName()
}

type AddressUpdate struct {
	UserID    uint   `form:"-"`
	No        string `form:"no"`
	Street    string `form:"street"`
	City      string `form:"city"`
	State     string `form:"state"`
	Country   string `form:"country"`
	Phone     string `form:"phone"`
	ZipCode   string `form:"zip_code"`
	IsDefault bool   `form:"is_default"` // Indicates whether this is the default address for the user
}

func (AddressUpdate) TableName() string {
	return Address{}.TableName()
}
