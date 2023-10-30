package model

import "new-mall/global"

type Product struct {
	global.Model
	ProductName          string `json:"productName" form:"productName" gorm:"column:product_name;type:varchar(200);"`
	ProductIntro         string `json:"productIntro" form:"productIntro" gorm:"column:product_intro;type:varchar(200);"`
	ProductCategoryId    int    `json:"productCategoryId" form:"productCategoryId" gorm:"column:product_category_id;type:bigint"`
	ProductCoverImg      string `json:"productCoverImg" form:"productCoverImg" gorm:"column:product_cover_img;type:varchar(200);"`
	ProductCarousel      string `json:"productCarousel" form:"productCarousel" gorm:"column:product_carousel;type:varchar(500);"`
	ProductDetailContent string `json:"productDetailContent" form:"productDetailContent" gorm:"column:product_detail_content;type:text;"`
	OriginalPrice        int    `json:"originalPrice" form:"originalPrice" gorm:"column:original_price;type:int"`
	SellingPrice         int    `json:"sellingPrice" form:"sellingPrice" gorm:"column:selling_price;type:int"`
	StockNum             int    `json:"stockNum" form:"stockNum" gorm:"column:stock_num;type:int"`
	Tag                  string `json:"tag" form:"tag" gorm:"column:tag;type:varchar(20);"`
	ProductSellStatus    int    `json:"productSellStatus" form:"productSellStatus" gorm:"column:product_sell_status;type:tinyint"`
	CreateUser           int    `json:"createUser" form:"createUser" gorm:"column:create_user;type:int"`
	UpdateUser           int    `json:"updateUser" form:"updateUser" gorm:"column:update_user;type:int"`
}

// TableName MallGoodsInfo
func (Product) TableName() string {
	return "products"
}
