package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/model"
	"new-mall/model/common"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, data *model.UserCreate) error {
	if err := r.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *userRepo) FindWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*model.User, error) {
	db := r.db

	var data model.User

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}

		return nil, err
	}

	return &data, nil
}

func (r *userRepo) ListWithCondition(
	ctx context.Context,
	filter *common.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.User, error) {
	var result []model.User
	db := r.db
	if v := filter.UserId; v > 0 {
		db = db.Where("id = ?", v)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(offset)

	if err := db.Table(model.User{}.TableName()).
		Count(&paging.Total).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}

func (r *userRepo) Update(
	ctx context.Context,
	id int,
	data *model.UserUpdate,
) error {
	if err := r.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *userRepo) Delete(ctx context.Context, id int) error {
	if err := r.db.Table(model.User{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"isDeleted": 1,
		}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
