package services

import (
	"context"
	"errors"
	"new-mall/internal/models"
	"new-mall/internal/repositories"
	"new-mall/internal/types"
	"new-mall/pkg/common"
)

type AddressService struct {
	AddressRepo *repositories.AddressRepository
}

func NewAddressService(addressRepo *repositories.AddressRepository) *AddressService {
	return &AddressService{
		AddressRepo: addressRepo,
	}
}

func (s *AddressService) CreateAddress(ctx context.Context, data *models.AddressCreate) error {
	if err := s.AddressRepo.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(models.AddressEntityName, err)
	}
	return nil
}

func (s *AddressService) GetAddress(ctx context.Context, id int) (*models.Address, error) {
	address, err := s.AddressRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(models.AddressEntityName, err)
	}

	return address, nil
}

func (s *AddressService) ListAddress(
	ctx context.Context,
	conditions map[string]interface{},
	filter *types.Filter,
	paging *types.Paging,
) ([]models.Address, error) {
	result, err := s.AddressRepo.ListWithCondition(ctx, conditions, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(models.AddressEntityName, err)
	}
	return result, nil
}

func (s *AddressService) DeleteAddress(ctx context.Context, id, userID uint) error {

	data, err := s.AddressRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrEntityNotFound(models.AddressEntityName, err)
		}
		return common.ErrCannotGetEntity(models.AddressEntityName, err)
	}

	if data.UserID != userID {
		return common.ErrCannotDeleteEntity(models.AddressEntityName, errors.New("not your address"))
	}

	if err = s.AddressRepo.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(models.AddressEntityName, err)
	}

	return nil
}

func (s *AddressService) UpdateAddress(ctx context.Context, id uint, data *models.AddressUpdate) error {
	oldData, err := s.AddressRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(models.AddressEntityName, err)
	}

	if oldData.UserID != data.UserID {
		return common.ErrCannotUpdateEntity(models.AddressEntityName, errors.New("not your address"))
	}

	if err = s.AddressRepo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
