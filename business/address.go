package business

import (
	"context"
	"errors"
	"new-mall/model"
	"new-mall/model/common"
)

type AddressRepo interface {
	Create(ctx context.Context, data *model.AddressCreate) error
	FindWithCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*model.Address, error)
	ListWithCondition(
		ctx context.Context,
		filter *common.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Address, error)
	Update(ctx context.Context, id int, data *model.AddressUpdate) error
	Delete(ctx context.Context, id int) error
}

type addressBiz struct {
	repo AddressRepo
}

func NewAddressBiz(repo AddressRepo) *addressBiz {
	return &addressBiz{repo: repo}
}

func (biz *addressBiz) CreateAddress(ctx context.Context, data *model.AddressCreate) error {
	err := biz.CreateAddress(ctx, data)
	return err
}

func (biz *addressBiz) FindAddress(ctx context.Context, id int) (*model.Address, error) {
	result, err := biz.repo.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if !errors.Is(err, common.RecordNotFound) {
			return nil, common.ErrEntityNotFound("Address", err)
		}
		return nil, common.ErrEntityNotFound("Address", err)
	}
	return result, nil
}

func (biz *addressBiz) ListAddress(
	ctx context.Context,
	filter *common.Filter,
	paging *common.Paging,
) ([]model.Address, error) {
	result, err := biz.repo.ListWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (biz *addressBiz) UpdateAddress(ctx context.Context, id int, data *model.AddressUpdate) error {
	oldData, err := biz.repo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.IsDeleted == 1 {
		return common.RecordNotFound
	}

	err = biz.repo.Update(ctx, id, data)
	return err
}

func (biz *addressBiz) DeleteAddress(ctx context.Context, id int) error {
	err := biz.repo.Delete(ctx, id)
	return err
}
