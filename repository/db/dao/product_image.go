package dao

import (
	"context"
	"gorm.io/gorm"
	"new-mall/repository/db/model"
	"new-mall/types"
)

type ProductImageDao struct {
	*gorm.DB
}

func NewProductImageDao(ctx context.Context) *ProductImageDao {
	return &ProductImageDao{NewDBClient(ctx)}
}

func NewProductImageDaoByDB(db *gorm.DB) *ProductImageDao {
	return &ProductImageDao{db}
}

func (dao *ProductImageDao) CreateProductImage(productImage *model.ProductImage) (err error) {
	err = dao.DB.Model(&model.ProductImage{}).Create(&productImage).Error

	return
}

func (dao *ProductImageDao) ListProductImageByProductId(pId uint) (r []*types.ProductImageRes, err error) {
	err = dao.DB.Model(&model.ProductImage{}).
		Where("product_id=?", pId).
		Find(&r).Error

	return
}
