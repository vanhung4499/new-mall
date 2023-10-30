package response

type ProductSearchResponse struct {
	ProductId       int    `json:"productId"`
	ProductName     string `json:"productName"`
	ProductIntro    string `json:"productIntro"`
	ProductCoverImg string `json:"productCoverImg"`
	SellingPrice    int    `json:"sellingPrice"`
}

type ProductInfoDetailResponse struct {
	ProductId            int      `json:"productId"`
	ProductName          string   `json:"productName"`
	ProductIntro         string   `json:"productIntro"`
	ProductCoverImg      string   `json:"productCoverImg"`
	SellingPrice         int      `json:"sellingPrice"`
	ProductDetailContent string   `json:"productDetailContent"  `
	OriginalPrice        int      `json:"originalPrice" `
	Tag                  string   `json:"tag" form:"tag" `
	ProductCarouselList  []string `json:"productCarouselList" `
}
