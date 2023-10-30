package request

import (
	"new-mall/model"
	"new-mall/model/common"
)

type ProductSearch struct {
	model.Product
	common.PageInfo
}

type ProductSearchParams struct {
	Keyword           string `form:"keyword"`
	ProductCategoryId int    `form:"productCategoryId"`
	OrderBy           string `form:"orderBy"`
	PageNumber        int    `form:"pageNumber"`
}

type ProductAddParam struct {
	ProductName          string `json:"productName"`
	ProductIntro         string `json:"productIntro"`
	ProductCategoryId    int    `json:"productCategoryId"`
	ProductCoverImg      string `json:"productCoverImg"`
	ProductCarousel      string `json:"productCarousel"`
	ProductDetailContent string `json:"productDetailContent"`
	OriginalPrice        string `json:"originalPrice"`
	SellingPrice         string `json:"sellingPrice"`
	StockNum             string `json:"stockNum"`
	Tag                  string `json:"tag"`
	ProductSellStatus    string `json:"productSellStatus"`
}

type ProductUpdateParam struct {
	ProductId            string `json:"productId"`
	ProductName          string `json:"productName"`
	ProductIntro         string `json:"productIntro"`
	ProductCategoryId    int    `json:"productCategoryId"`
	ProductCoverImg      string `json:"productCoverImg"`
	ProductCarousel      string `json:"productCarousel"`
	ProductDetailContent string `json:"productDetailContent"`
	OriginalPrice        string `json:"originalPrice"`
	SellingPrice         int    `json:"sellingPrice"`
	StockNum             string `json:"stockNum"`
	Tag                  string `json:"tag"`
	ProductSellStatus    int    `json:"productSellStatus"`
	UpdateUser           int    `json:"updateUser" form:"updateUser" gorm:"column:update_user;type:int"`
}

type StockNumDTO struct {
	ProductId    int `json:"productId"`
	ProductCount int `json:"productCount"`
}
