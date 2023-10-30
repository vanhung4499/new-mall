package business

import (
	"context"
	"new-mall/model"
	"new-mall/model/common"
	"new-mall/repository"
)

type UserRepo interface {
	CreateUser(ctx context.Context, data *model.UserCreate) error
}

func NewUserBiz(repo repository.UserRepo) *userBiz {
	return &userBiz{repo: repo}
}

type userBiz struct {
	repo repository.UserRepo
}

func (biz *userBiz) CreateUser(ctx context.Context, data *model.UserCreate) error {
	err := biz.CreateUser(ctx, data)
	return err
}

func (biz *userBiz) UpdateUser(ctx context.Context, id int, data *model.UserUpdate) error {
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

func (biz *userBiz) DeleteUser(ctx context.Context, id int) error {
	err := biz.repo.Delete(ctx, id)
	return err
}
