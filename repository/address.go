package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/model"
	"new-mall/model/common"
)

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) *addressRepo {
	return &addressRepo{db: db}
}

func (r *addressRepo) FindWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*model.Address, error) {
	db := r.db

	var data model.Address

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

func (r *addressRepo) ListWithCondition(
	ctx context.Context,
	filter *common.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Address, error) {
	var result []model.Address
	db := r.db
	if v := filter.UserId; v > 0 {
		db = db.Where("user_id = ?", v)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(offset)

	if err := db.Table(model.Address{}.TableName()).
		Count(&paging.Total).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}

func (r *addressRepo) Create(ctx context.Context, data *model.AddressCreate) error {
	if err := r.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *addressRepo) Update(ctx context.Context, id int, data *model.AddressUpdate) error {
	if err := r.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *addressRepo) Delete(ctx context.Context, id int) error {
	if err := r.db.Table(model.Address{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"isDeleted": 1,
		}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
