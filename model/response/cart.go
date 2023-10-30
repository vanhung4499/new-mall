package response

type CartItemResponse struct {
	CartItemId int `json:"cartItemId"`

	ProductId int `json:"productId"`

	ProductCount int `json:"productCount"`

	ProductName string `json:"productName"`

	ProductCoverImg string `json:"productCoverImg"`

	SellingPrice int `json:"sellingPrice"`
}
