package services

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/internal/models"
	"new-mall/internal/repositories"
	"new-mall/pkg/common"
)

type CartService struct {
	CartRepo    *repositories.CartRepository
	ProductRepo *repositories.ProductRepository
}

func NewCartService(cartRepo *repositories.CartRepository, productRepo *repositories.ProductRepository) *CartService {
	return &CartService{
		CartRepo:    cartRepo,
		ProductRepo: productRepo,
	}
}

func (s *CartService) GetCart(ctx context.Context, userID uint) (*models.Cart, error) {
	condition := map[string]interface{}{"user_id": userID}
	cart, err := s.CartRepo.FindCartWithCondition(ctx, condition, "CartItems", "CartItems.Product", "CartItems.Product.Images")
	if err != nil {
		return nil, common.ErrEntityNotFound(models.CartEntityName, err)
	}

	return cart, nil
}

// CreateCart creates a new shopping cart for a user
func (s *CartService) CreateCart(ctx context.Context, data *models.CartCreate) error {

	// Check if the user has an existing cart
	_, err := s.CartRepo.FindCartWithCondition(ctx, map[string]interface{}{"user_id": data.UserID})
	if errors.Is(err, common.RecordNotFound) {
		// If the user doesn't have a cart, create a new one
		if err = s.CartRepo.Create(data); err != nil {
			return common.ErrCannotCreateEntity(models.CartEntityName, err)
		}
	} else if err != nil {
		return common.ErrCannotGetEntity(models.CartEntityName, err)
	}

	return nil
}

func (s *CartService) AddToCart(ctx context.Context, userID uint, data *models.CartItemCreate) error {

	// Check if the product exists
	product, err := s.ProductRepo.FindWithCondition(ctx, map[string]interface{}{"id": data.ProductID})
	if err != nil {
		return common.ErrEntityNotFound(models.ProductEntityName, err)
	}

	// Check if the user has an existing cart
	cart, err := s.CartRepo.FindCartWithCondition(ctx, map[string]interface{}{"user_id": userID})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// If the user doesn't have a cart, create a new one
		newCart := models.CartCreate{UserID: userID}

		if err = s.CartRepo.Create(&newCart); err != nil {
			return common.ErrCannotCreateEntity(models.CartEntityName, err)
		}
	} else if err != nil {
		return common.ErrCannotGetEntity(models.CartEntityName, err)
	}

	// Check if the product is already in the cart
	var cartItem *models.CartItem

	condition := map[string]interface{}{"cart_id": cart.ID, "product_id": product.ID}

	if cartItem, err = s.CartRepo.FindCartItemWithCondition(ctx, condition); err == nil {
		// If the product is already in the cart, update the quantity
		cartItem.Quantity += data.Quantity

		if err = s.CartRepo.UpdateCartItem(ctx, cartItem); err != nil {
			return common.ErrCannotUpdateEntity(models.CartItemEntityName, err)
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// If the product is not in the cart, add a new cart item
		if err = s.CartRepo.CreateCartItem(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(models.CartItemEntityName, err)
		}
	} else {
		return common.ErrCannotGetEntity(models.CartItemEntityName, err)
	}

	// Recalculate the total amount of the cart
	if err = s.CartRepo.RecalculateCartTotal(cartItem.CartID); err != nil {
		return common.ErrCannotUpdateEntity(models.CartEntityName, err)
	}

	return nil
}

// DeleteCart delete the entire shopping cart for a user
func (s *CartService) DeleteCart(ctx context.Context, userID uint) error {

	// Retrieve the user's shopping cart
	cart, err := s.CartRepo.FindCartWithCondition(ctx, map[string]interface{}{"user_id": userID})
	if err != nil {
		return common.ErrEntityNotFound(models.CartEntityName, err)
	}

	// Delete all cart items associated with the cart
	err = s.CartRepo.DeleteCartItemWithCondition(ctx, map[string]interface{}{"cart_id": cart.ID})
	if err != nil {
		return common.ErrCannotDeleteEntity(models.CartItemEntityName, err)
	}

	// Delete the cart
	err = s.CartRepo.DeleteCart(ctx, cart.ID)
	if err != nil {
		return common.ErrCannotDeleteEntity(models.CartEntityName, err)
	}

	return nil
}

// UpdateCart updates the quantity of a product in the user's shopping cart
func (s *CartService) UpdateCart(ctx context.Context, userID uint, data *models.CartItemUpdate) error {

	// Check if the user has a cart
	cart, err := s.CartRepo.FindCartWithCondition(ctx, map[string]interface{}{"user_id": userID})
	if errors.Is(err, common.RecordNotFound) {
		return common.ErrEntityNotFound(models.CartEntityName, err)
	} else if err != nil {
		return common.ErrCannotGetEntity(models.CartEntityName, err)
	}

	// Check if the product is in the cart
	condition := map[string]interface{}{"cart_id": cart.ID, "product_id": data.ProductID}
	cartItem, err := s.CartRepo.FindCartItemWithCondition(ctx, condition)
	if errors.Is(err, common.RecordNotFound) {
		return common.ErrEntityNotFound(models.CartItemEntityName, err)
	} else if err != nil {
		return common.ErrCannotGetEntity(models.CartItemEntityName, err)
	}

	// Update the quantity and price
	cartItem.Quantity = data.Quantity
	cartItem.Price = data.Price
	if err = s.CartRepo.UpdateCartItem(ctx, cartItem); err != nil {
		return common.ErrCannotUpdateEntity(models.CartItemEntityName, err)
	}

	// Recalculate the total amount of the cart
	if err = s.CartRepo.RecalculateCartTotal(cartItem.CartID); err != nil {
		return common.ErrCannotUpdateEntity(models.CartEntityName, err)
	}

	return nil
}

func (s *CartService) DeleteCartItemWithCondition(ctx context.Context, userID uint, productId uint) error {

	// Retrieve the user's shopping cart
	cart, err := s.CartRepo.FindCartWithCondition(ctx, map[string]interface{}{"user_id": userID})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrEntityNotFound(models.CartEntityName, err)
		}
		return common.ErrCannotGetEntity(models.CartEntityName, err)
	}

	// Delete cart item associated with the cart
	condition := map[string]interface{}{"cart_id": cart.ID, "product_id": productId}
	err = s.CartRepo.DeleteCartItemWithCondition(ctx, condition)
	if err != nil {
		return common.ErrCannotDeleteEntity(models.CartItemEntityName, err)
	}

	// Recalculate the total amount of the cart
	if err = s.CartRepo.RecalculateCartTotal(cart.ID); err != nil {
		return common.ErrCannotUpdateEntity(models.CartEntityName, err)
	}

	return nil
}
