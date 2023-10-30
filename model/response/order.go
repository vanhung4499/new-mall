package response

import (
	"time"
)

type OrderResponse struct {
	OrderId           int           `json:"orderId"`
	OrderNo           string        `json:"orderNo"`
	TotalPrice        int           `json:"totalPrice"`
	PayType           int           `json:"payType"`
	OrderStatus       int           `json:"orderStatus"`
	OrderStatusString string        `json:"orderStatusString"`
	CreateAt          time.Time     `json:"createAt"`
	OrderItemVOS      []OrderItemVO `json:"orderItemVOS"`
}

type OrderItemVO struct {
	ProductId       int    `json:"productId"`
	ProductName     string `json:"productName"`
	ProductCount    int    `json:"productCount"`
	ProductCoverImg string `json:"productCoverImg"`
	SellingPrice    int    `json:"sellingPrice"`
}

type OrderDetailVO struct {
	OrderNo           string        `json:"orderNo"`
	TotalPrice        int           `json:"totalPrice"`
	PayStatus         int           `json:"payStatus"`
	PayType           int           `json:"payType"`
	PayTypeString     string        `json:"payTypeString"`
	PayAt             time.Time     `json:"payAt"`
	OrderStatus       int           `json:"orderStatus"`
	OrderStatusString string        `json:"orderStatusString"`
	CreateAt          time.Time     `json:"createAt"`
	OrderItemVOS      []OrderItemVO `json:"orderItemVOS"`
}
