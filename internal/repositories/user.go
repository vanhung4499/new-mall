package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Update(ctx context.Context, id uint, data *models.UserUpdate) error {

	if err := r.DB.
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (r *UserRepository) Save(ctx context.Context, user *models.User) error {
	if err := r.DB.
		Model(&models.User{}).
		Save(user).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (r *UserRepository) Create(ctx context.Context, data *models.UserCreate) error {
	db := r.DB
	if err := db.
		Table(models.UserCreate{}.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	if err := r.DB.
		Where("id = ?", id).
		Delete(&models.User{}, id).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *UserRepository) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*models.User, error) {
	var result models.User

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

func (r *UserRepository) ListWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
	sorting *types.Sorting,
	moreKeys ...string,
) ([]models.User, error) {

	db := r.DB

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(offset)

	var result []models.User

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
