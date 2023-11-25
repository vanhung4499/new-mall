package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) FindWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*models.Product, error) {
	db := r.DB

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	var product models.Product

	if err := db.
		Model(&models.Product{}).
		Where(conditions).
		First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrDB(err)
		}
		return nil, common.ErrDB(err)
	}

	return &product, nil
}

func (r *ProductRepository) ListWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
	moreKeys ...string,
) ([]models.Product, error) {
	db := r.DB
	db = db.Table(models.Product{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.CategoryID != 0 {
			db = db.Where("category_id = ?", v.CategoryID)
		}
		if v.Keyword != "" {
			db = db.Where("name LIKE ? AND description LIKE ?", "%"+v.Keyword+"%", "%"+v.Keyword+"%")
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(offset)

	var result []models.Product

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (r *ProductRepository) Create(ctx context.Context, product *models.ProductCreate) error {
	if err := r.DB.
		Table(models.ProductCreate{}.TableName()).
		Create(&product).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id int) error {
	if err := r.DB.
		Table(models.Product{}.TableName()).
		Where("id = ?", id).
		Delete(&models.Product{}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *ProductRepository) Update(ctx context.Context, id int, product *models.ProductUpdate) error {
	if err := r.DB.
		Table(models.ProductCreate{}.TableName()).
		Where("id = ?", id).
		Updates(&product).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
