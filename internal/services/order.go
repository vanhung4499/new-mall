package services

import (
	"context"
	"new-mall/internal/models"
	"new-mall/internal/repositories"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type OrderService struct {
	OrderRepo   *repositories.OrderRepository
	AddressRepo *repositories.AddressRepository
	CartRepo    *repositories.CartRepository
	UserRepo    *repositories.UserRepository
}

func NewOrderService(
	orderRepo *repositories.OrderRepository,
	addressRepo *repositories.AddressRepository,
	cartRepo *repositories.CartRepository,
	userRepo *repositories.UserRepository,
) *OrderService {
	return &OrderService{
		OrderRepo:   orderRepo,
		AddressRepo: addressRepo,
		CartRepo:    cartRepo,
		UserRepo:    userRepo,
	}
}

func (s *OrderService) CreateOrderFromCart(ctx context.Context, userID, addressID uint) error {
	// Find cart
	condition := map[string]interface{}{"user_id": userID}
	cart, err := s.CartRepo.FindCartWithCondition(ctx, condition, "CartItems")
	if err != nil {
		return common.ErrEntityNotFound(models.CartEntityName, err)
	}

	cartItems := cart.CartItems

	// Check if cart is empty
	if cartItems == nil || len(cartItems) == 0 {
		return models.ErrCartIsEmpty
	}

	// Create an order
	order := &models.Order{
		UserID:      userID,
		OrderItems:  make([]models.OrderItem, len(cartItems)),
		TotalAmount: cart.TotalAmount,
		AddressID:   addressID,
		Status:      1,
	}

	// Populate order items
	for _, cartItem := range cartItems {
		orderItem := models.OrderItem{
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
			Price:     cartItem.Price,
		}
		order.OrderItems = append(order.OrderItems, orderItem)
	}

	// Save order to database
	if err = s.OrderRepo.CreateOrder(ctx, order); err != nil {
		return common.ErrCannotCreateEntity(models.OrderEntityName, err)
	}

	// Delete cart
	if err = s.CartRepo.DeleteCart(ctx, cart.ID); err != nil {
		return common.ErrCannotDeleteEntity(models.CartEntityName, err)
	}

	return nil
}

func (s *OrderService) ListOrder(ctx context.Context, userID uint, paging *types.Paging) ([]models.Order, error) {
	var orders []models.Order
	orders, err := s.OrderRepo.ListWithCondition(
		ctx,
		map[string]interface{}{"user_id": userID},
		nil,
		paging,
		"OrderItems", "OrderItems.Product", "OrderItems.Product.Images")
	if err != nil {
		return nil, common.ErrCannotListEntity(models.OrderEntityName, err)
	}

	return orders, nil
}

func (s *OrderService) GetOrder(ctx context.Context, userID, id uint) (*models.Order, error) {

	order, err := s.OrderRepo.FindWithCondition(ctx, map[string]interface{}{"id": id}, "OrderItems", "OrderItems.Product", "OrderItems.Product.Images")
	if err != nil {
		return nil, common.ErrEntityNotFound(models.OrderEntityName, err)
	}

	if order.UserID != userID {
		return nil, common.ErrNoPermission(err)
	}

	return order, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, userID, id uint) error {
	order, err := s.OrderRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(models.OrderEntityName, err)
	}

	if order.UserID != userID {
		return common.ErrNoPermission(err)
	}

	if err = s.OrderRepo.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(models.OrderEntityName, err)
	}

	return nil
}
