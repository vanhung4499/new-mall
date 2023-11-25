package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type ImageRepository struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db}
}

func (r *ImageRepository) FindWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*models.Image, error) {
	db := r.DB

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	var image models.Image

	if err := db.
		Model(&models.Image{}).
		Where(conditions).
		First(&image).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrDB(err)
		}
		return nil, common.ErrDB(err)
	}

	return &image, nil
}

func (r *ImageRepository) ListWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
	moreKeys ...string,
) ([]models.Image, error) {
	db := r.DB

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	var result []models.Image
	if err := db.Where(conditions).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (r *ImageRepository) Create(ctx context.Context, data *models.ImageCreate) error {
	if err := r.DB.
		Table(models.ImageCreate{}.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *ImageRepository) Delete(ctx context.Context, id uint) error {
	if err := r.DB.
		Where("id = ?", id).
		Delete(&models.Image{}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
