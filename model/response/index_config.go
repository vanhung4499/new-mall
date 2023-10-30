package response

type IndexConfigProductResponse struct {
	ProductId       int    `json:"productId"`
	ProductName     string `json:"productName"`
	ProductIntro    string `json:"productIntro"`
	ProductCoverImg string `json:"productCoverImg"`
	SellingPrice    int    `json:"sellingPrice"`
	Tag             string `json:"tag"`
}
