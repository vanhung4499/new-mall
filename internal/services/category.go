package services

import (
	"context"
	"new-mall/internal/models"
	"new-mall/internal/repositories"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type CategoryService struct {
	CategoryRepo *repositories.CategoryRepository
}

func NewCategoryService(categoryRepo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepo: categoryRepo,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, data *models.CategoryCreate) error {
	if err := s.CategoryRepo.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(models.CategoryEntityName, err)
	}
	return nil
}

func (s *CategoryService) GetCategory(ctx context.Context, id int) (*models.Category, error) {

	result, err := s.CategoryRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(models.CategoryEntityName, err)
	}

	return result, nil
}

func (s *CategoryService) ListCategory(
	ctx context.Context,
	filter *types.Filter,
	paging *types.Paging,
) ([]models.Category, error) {
	result, err := s.CategoryRepo.ListWithCondition(ctx, nil, nil, nil)
	if err != nil {
		return nil, common.ErrCannotListEntity(models.CategoryEntityName, err)
	}
	return result, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id int) error {

	data, err := s.CategoryRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if data == nil {
		return common.ErrEntityNotFound(models.CategoryEntityName, err)
	}

	if err = s.CategoryRepo.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(models.CategoryEntityName, err)
	}

	return nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id int, data *models.CategoryUpdate) error {
	_, err := s.CategoryRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(models.CategoryEntityName, err)
	}

	if err = s.CategoryRepo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
