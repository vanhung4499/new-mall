package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		DB: db,
	}
}

func (r *AddressRepository) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*models.Address, error) {
	var result models.Address

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

func (r *AddressRepository) ListWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
	moreKeys ...string,
) ([]models.Address, error) {
	db := r.DB
	db = db.Where(conditions)

	if v := filter; v != nil {
		if v.UserID != 0 {
			db = db.Where("user_id = ?", v.UserID)
		}
	}

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	var result []models.Address
	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (r *AddressRepository) Create(ctx context.Context, data *models.AddressCreate) error {
	if err := r.DB.
		Table(models.AddressCreate{}.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *AddressRepository) Delete(ctx context.Context, id uint) error {
	if err := r.DB.
		Where("id = ?", id).
		Delete(&models.Address{}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *AddressRepository) Update(ctx context.Context, id uint, data *models.AddressUpdate) error {
	if err := r.DB.
		Table(models.AddressUpdate{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
