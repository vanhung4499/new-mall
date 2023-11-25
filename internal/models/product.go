package models

import (
	"gorm.io/gorm"
)

const ProductEntityName = "Product"

type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	Description   string `gorm:"size:1000"`
	ImageURL      string `gorm:"not null"`
	Price         float64
	DiscountPrice float64
	OnSale        bool `gorm:"default:false"`
	CategoryID    uint `gorm:"not null"`
	Category      Category
	StockQuantity int
	Images        []Image
}

func (Product) TableName() string {
	return "products"
}

type ProductCreate struct {
	gorm.Model
	Name          string        `form:"name"`
	Description   string        `form:"description"`
	ImageURL      string        `form:"image_url"`
	Price         float64       `form:"price"`
	DiscountPrice float64       `form:"discount_price"`
	OnSale        bool          `form:"on_sale"`
	CategoryID    uint          `form:"category_id"`
	StockQuantity int           `form:"stock_quantity"`
	Images        []ImageCreate `form:"images,omitempty"`
}

func (ProductCreate) TableName() string {
	return Product{}.TableName()
}

type ProductUpdate struct {
	Name          string        `form:"name"`
	Description   string        `form:"description"`
	ImageURL      string        `form:"image_url"`
	Price         float64       `form:"price"`
	DiscountPrice float64       `form:"discount_price"`
	OnSale        bool          `form:"on_sale"`
	CategoryID    uint          `form:"category_id"`
	StockQuantity int           `form:"stock_quantity"`
	Images        []ImageCreate `form:"images,omitempty"`
}

func (ProductUpdate) TableName() string {
	return Product{}.TableName()
}
