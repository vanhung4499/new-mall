package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type OrderRepository struct {
	*gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *models.OrderCreate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return common.ErrDB(err)
		}

		for _, orderItem := range order.OrderItems {
			orderItem.OrderID = order.ID
			if err := tx.Create(&orderItem).Error; err != nil {
				return common.ErrDB(err)
			}
		}

		return nil
	})
}

func (r *OrderRepository) ListWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
	moreKeys ...string,
) ([]models.Order, error) {
	db := r.DB
	db = db.
		Table(models.Order{}.TableName()).
		Where(condition)

	if v := filter; v != nil {
		if v.UserID > 0 {
			db = db.Where("user_id = ?", v.UserID)
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

	var orders []models.Order

	err := db.
		Limit(paging.Limit).
		Order("id DESC").
		Find(&orders).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return orders, nil
}

// FindWithCondition Get order details
func (r *OrderRepository) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*models.Order, error) {
	db := r.DB

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	var order models.Order

	if err := db.
		Where(condition).
		First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &order, nil
}

func (r *OrderRepository) Delete(ctx context.Context, id uint) error {
	if err := r.DB.
		Model(&models.Order{}).
		Where("id = ?", id).
		Delete(&models.Order{}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
