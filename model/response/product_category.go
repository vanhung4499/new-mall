package response

import (
	"new-mall/model"
)

type ProductCategoryResponse struct {
	ProductCategory model.Category `json:"productCategory"`
}
