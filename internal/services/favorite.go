package services

import (
	"context"
	"new-mall/internal/models"
	"new-mall/internal/repositories"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type FavoriteService struct {
	FavoriteRepo *repositories.FavoriteRepository
	UserRepo     *repositories.UserRepository
	ProductRepo  *repositories.ProductRepository
}

func NewFavoriteService(favoriteRepo *repositories.FavoriteRepository, userRepo *repositories.UserRepository, productRepo *repositories.ProductRepository) *FavoriteService {
	return &FavoriteService{
		FavoriteRepo: favoriteRepo,
		UserRepo:     userRepo,
		ProductRepo:  productRepo,
	}
}

func (s *FavoriteService) CreateFavorite(ctx context.Context, data *models.FavoriteCreate) error {
	_, err := s.FavoriteRepo.FindWithCondition(ctx, map[string]interface{}{"user_id": data.UserID, "product_id": data.ProductID})

	if err == nil {
		return common.ErrEntityAlreadyExisted(models.FavoriteEntityName, err)
	}

	if err = s.FavoriteRepo.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(models.FavoriteEntityName, err)
	}
	return nil
}

func (s *FavoriteService) GetFavorite(ctx context.Context, id int) (*models.Favorite, error) {

	address, err := s.FavoriteRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(models.FavoriteEntityName, err)
	}

	return address, nil
}

func (s *FavoriteService) ListFavorite(
	ctx context.Context,
	condition map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
) ([]models.Favorite, error) {
	result, err := s.FavoriteRepo.ListWithCondition(ctx, condition, filter, paging, "Product", "Product.Images")
	if err != nil {
		return nil, common.ErrCannotListEntity(models.FavoriteEntityName, err)
	}
	return result, nil
}

func (s *FavoriteService) DeleteFavorite(ctx context.Context, userId, productID uint) error {

	data, err := s.FavoriteRepo.FindWithCondition(ctx, map[string]interface{}{"user_id": userId, "product_id": productID})

	if data == nil {
		return common.ErrEntityNotFound(models.FavoriteEntityName, err)
	}

	if err = s.FavoriteRepo.Delete(ctx, data.ID); err != nil {
		return common.ErrCannotDeleteEntity(models.FavoriteEntityName, err)
	}

	return nil
}
