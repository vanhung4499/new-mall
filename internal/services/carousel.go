package services

import (
	"context"
	"new-mall/internal/models"
	"new-mall/internal/repositories"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type CarouselService struct {
	CarouselRepo *repositories.CarouselRepository
}

func NewCarouselService(carouselRepo *repositories.CarouselRepository) *CarouselService {
	return &CarouselService{
		CarouselRepo: carouselRepo,
	}
}

func (s *CarouselService) CreateCarousel(ctx context.Context, data *models.CarouselCreate) error {

	// TODO: Check if the user has permission to create this carousel

	if err := s.CarouselRepo.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(models.CarouselEntityName, err)
	}

	return nil
}

func (s *CarouselService) GetCarousel(ctx context.Context, id int) (*models.Carousel, error) {
	result, err := s.CarouselRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(models.CarouselEntityName, err)
	}

	return result, nil
}

func (s *CarouselService) ListCarousel(
	ctx context.Context,
	filter *types.Filter,
	paging *types.Paging,
) ([]models.Carousel, error) {
	result, err := s.CarouselRepo.ListWithCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(models.CarouselEntityName, err)
	}
	return result, nil
}

func (s *CarouselService) DeleteCarousel(ctx context.Context, id int) error {

	data, err := s.CarouselRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if data == nil {
		return common.ErrEntityNotFound(models.CarouselEntityName, err)
	}

	if err = s.CarouselRepo.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(models.CarouselEntityName, err)
	}

	return nil
}

func (s *CarouselService) UpdateCarousel(ctx context.Context, id int, data *models.CarouselUpdate) error {
	_, err := s.CarouselRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(models.CarouselEntityName, err)
	}

	if err = s.CarouselRepo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
