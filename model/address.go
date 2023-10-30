package model

import "new-mall/global"

// Address Structure
type Address struct {
	global.Model
	UserId        int    `json:"userId" form:"userId" gorm:"column:user_id;type:bigint"`
	UserName      string `json:"userName" form:"userName" gorm:"column:user_name;type:varchar(30);"`
	UserPhone     string `json:"userPhone" form:"userPhone" gorm:"column:user_phone;type:varchar(11);"`
	DefaultFlag   int    `json:"defaultFlag" form:"defaultFlag" gorm:"column:default_flag;type:tinyint"`
	ProvinceName  string `json:"provinceName" form:"provinceName" gorm:"column:province_name;type:varchar(32);"`
	CityName      string `json:"cityName" form:"cityName" gorm:"column:city_name;type:varchar(32);"`
	RegionName    string `json:"regionName" form:"regionName" gorm:"column:region_name;type:varchar(32);"`
	DetailAddress string `json:"detailAddress" form:"detailAddress" gorm:"column:detail_address;type:varchar(64);"`
}

// TableName Address
func (Address) TableName() string {
	return "addresses"
}

type AddressCreate struct {
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	DefaultFlag   byte   `json:"defaultFlag"` // 0-no 1-yes
	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
}

// TableName Address
func (AddressCreate) TableName() string {
	return "addresses"
}

type AddressUpdate struct {
	UserId        int    `json:"userId"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	DefaultFlag   byte   `json:"defaultFlag"` // 0-no 1-yes
	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
}

// TableName Address
func (AddressUpdate) TableName() string {
	return "addresses"
}
