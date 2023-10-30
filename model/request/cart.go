package request

import (
	"new-mall/model/common"
)

type CartSearch struct {
	common.PageInfo
}

type SaveCartItemParam struct {
	ProductCount int `json:"productCount"`
	ProductId    int `json:"productId"`
}

type UpdateCartItemParam struct {
	CartItemId   int `json:"cartItemId"`
	ProductCount int `json:"productCount"`
}
