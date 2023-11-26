package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		DB: db,
	}
}

func (r *CartRepository) FindCartWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*models.Cart, error) {
	db := r.DB
	db = db.
		Table(models.Cart{}.TableName()).
		Where(condition)

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	var cart models.Cart

	if err := db.
		First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &cart, nil
}

// Create creates a new shopping cart in the database
func (r *CartRepository) Create(cart *models.CartCreate) error {
	if err := r.DB.
		Table(models.CartCreate{}.TableName()).
		Create(cart).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

// DeleteCart deletes a user's shopping cart from the database
func (r *CartRepository) DeleteCart(ctx context.Context, id uint) error {
	if err := r.DB.Delete(&models.Cart{}, id).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

// CreateCartItem creates a new cart item in the database
func (r *CartRepository) CreateCartItem(ctx context.Context, cartItem *models.CartItemCreate) error {
	if err := r.DB.
		Table(models.CartItemCreate{}.TableName()).
		Create(cartItem).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

// UpdateCartItem updates a cart item in the database
func (r *CartRepository) UpdateCartItem(ctx context.Context, cartItem *models.CartItem) error {
	if err := r.DB.
		Save(cartItem).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *CartRepository) DeleteCartItem(ctx context.Context, id uint) error {
	if err := r.DB.
		Table(models.CartItem{}.TableName()).
		Where("id = ?", id).
		Delete(&models.CartItem{}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (r *CartRepository) FindCartItemWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*models.CartItem, error) {
	db := r.DB

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	var cartItem models.CartItem

	if err := db.
		Where(condition).
		First(&cartItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &cartItem, nil
}

func (r *CartRepository) ListCartItemWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
	moreKeys ...string,
) ([]models.CartItem, error) {
	db := r.DB
	db = db.
		Table(models.CartItem{}.TableName()).
		Where(condition)

	var result []models.CartItem
	if err := r.DB.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (r *CartRepository) DeleteCartItemWithCondition(
	ctx context.Context,
	condition map[string]interface{},
) error {
	if err := r.DB.
		Model(models.CartItem{}).
		Where(condition).
		Delete(&models.CartItem{}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

// RecalculateCartTotal recalculates the total amount of a cart based on the cart items
func (r *CartRepository) RecalculateCartTotal(cartID uint) error {

	// Retrieve the cart items associated with the cart
	var cartItems []models.CartItem
	if err := r.DB.
		Model(&models.CartItem{}).
		Where("cart_id = ?", cartID).
		Find(&cartItems).Error; err != nil {
		return common.ErrDB(err)
	}

	totalAmount := 0.0
	// Calculate the total amount
	for _, cartItem := range cartItems {
		totalAmount += cartItem.Price * float64(cartItem.Quantity)
	}

	// Update the total amount of the cart in the database
	if err := r.DB.
		Model(&models.Cart{}).
		Where("id = ?", cartID).
		Update("total_amount", totalAmount).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
