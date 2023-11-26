package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type FavoriteRepository struct {
	DB *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{
		DB: db,
	}
}

func (r *FavoriteRepository) Create(ctx context.Context, data *models.FavoriteCreate) error {
	if err := r.DB.
		Table(models.Favorite{}.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (r *FavoriteRepository) Delete(ctx context.Context, id uint) error {
	if err := r.DB.
		Where("id = ?", id).
		Delete(&models.Favorite{}, id).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *FavoriteRepository) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*models.Favorite, error) {
	var result models.Favorite

	db := r.DB
	db = db.Where(condition)

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	if err := db.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}

func (r *FavoriteRepository) ListWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
	moreKeys ...string,
) ([]models.Favorite, error) {
	db := r.DB
	db = db.Where(condition)

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(offset)

	var result []models.Favorite

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
