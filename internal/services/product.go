package services

import (
	"context"
	"new-mall/internal/models"
	"new-mall/internal/repositories"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type ProductService struct {
	ProductRepo *repositories.ProductRepository
	ImageRepo   *repositories.ImageRepository
}

func NewProductService(productRepo *repositories.ProductRepository, imageRepo *repositories.ImageRepository) *ProductService {
	return &ProductService{
		ProductRepo: productRepo,
		ImageRepo:   imageRepo,
	}
}

func (s *ProductService) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	result, err := s.ProductRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(models.ProductEntityName, err)
	}

	return result, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, data *models.ProductCreate) error {
	if err := s.ProductRepo.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(models.ProductEntityName, err)
	}
	return nil
}

func (s *ProductService) ListProduct(
	ctx context.Context,
	filter *types.Filter,
	paging *types.Paging,
) ([]models.Product, error) {

	result, err := s.ProductRepo.ListWithCondition(ctx, nil, filter, paging, "Category", "Images")
	if err != nil {
		return nil, common.ErrCannotListEntity(models.ProductEntityName, err)
	}

	return result, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	data, err := s.ProductRepo.FindWithCondition(ctx, map[string]interface{}{"id": id}, "Images")
	if err != nil {
		return common.ErrEntityNotFound(models.ProductEntityName, err)
	}

	for _, image := range data.Images {
		if err = s.ImageRepo.Delete(ctx, image.ID); err != nil {
			return common.ErrCannotDeleteEntity(models.ImageEntityName, err)
		}
	}

	if err = s.ProductRepo.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(models.ProductEntityName, err)
	}

	return nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, data *models.ProductUpdate) error {
	_, err := s.ProductRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(models.ProductEntityName, err)
	}

	if err = s.ProductRepo.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(models.ProductEntityName, err)
	}

	return nil
}

func (s *ProductService) ListProductImage(ctx context.Context, productID int) ([]string, error) {
	result, err := s.ImageRepo.ListWithCondition(ctx, map[string]interface{}{"product_id": productID}, nil, nil)
	if err != nil {
		return nil, err
	}

	imageList := make([]string, len(result))
	for i, image := range result {
		imageList[i] = image.ImageURL
	}

	return imageList, nil
}
