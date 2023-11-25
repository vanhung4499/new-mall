package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (r *CategoryRepository) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*models.Category, error) {
	var result models.Category

	db := r.DB

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	if err := db.Where(condition).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound

		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}

func (r *CategoryRepository) ListWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
	moreKeys ...string,
) ([]models.Category, error) {
	db := r.DB

	if v := filter; v != nil {
		if v.UserID != 0 {
			db = db.Where("user_id = ?", v.UserID)
		}
	}

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	var result []models.Category
	if err := db.Where(condition).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (r *CategoryRepository) Create(ctx context.Context, data *models.CategoryCreate) error {
	if err := r.DB.
		Table(models.CategoryCreate{}.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id int) error {
	if err := r.DB.
		Where("id = ?", id).
		Delete(&models.Category{}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *CategoryRepository) Update(ctx context.Context, id int, data *models.CategoryUpdate) error {
	if err := r.DB.
		Table(models.CategoryUpdate{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
