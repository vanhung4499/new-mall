package types

type ProductImageRes struct {
	ProductID uint   `json:"product_id" form:"product_id"`
	ImagePath string `json:"image_path" form:"image_path"`
}
