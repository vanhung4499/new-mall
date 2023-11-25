package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type CarouselRepository struct {
	DB *gorm.DB
}

func NewCarouselRepository(db *gorm.DB) *CarouselRepository {
	return &CarouselRepository{
		DB: db,
	}
}

func (r *CarouselRepository) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*models.Carousel, error) {
	var result models.Carousel

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

func (r *CarouselRepository) ListWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
	moreKeys ...string,
) ([]models.Carousel, error) {
	db := r.DB

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	var result []models.Carousel
	if err := db.Where(condition).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (r *CarouselRepository) Create(ctx context.Context, data *models.CarouselCreate) error {
	if err := r.DB.
		Table(models.CarouselCreate{}.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *CarouselRepository) Delete(ctx context.Context, id int) error {
	if err := r.DB.
		Where("id = ?", id).
		Delete(&models.Carousel{}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *CarouselRepository) Update(ctx context.Context, id int, data *models.CarouselUpdate) error {
	if err := r.DB.
		Table(models.CarouselUpdate{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
