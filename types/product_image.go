package types

type ProductImageRes struct {
	ProductID uint   `json:"product_id" form:"product_id"`
	ImagePath string `json:"img_path" form:"img_path"`
}
