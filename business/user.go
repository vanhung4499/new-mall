package business

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"new-mall/model"
	"new-mall/model/common"
	"new-mall/model/request"
	"new-mall/utils/jwt"
)

type UserRepo interface {
	Create(ctx context.Context, data *model.UserCreate) error
	FindWithCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*model.User, error)
	ListWithCondition(
		ctx context.Context,
		filter *common.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.User, error)
	Update(ctx context.Context, id int, data *model.UserUpdate) error
	Delete(ctx context.Context, id int) error
}

func NewUserBiz(repo UserRepo) *userBiz {
	return &userBiz{repo: repo}
}

type userBiz struct {
	repo UserRepo
}

func (biz *userBiz) LoginUser(ctx context.Context, req *request.LoginUserRequest) (string, error) {
	user, err := biz.repo.FindWithCondition(ctx, map[string]interface{}{"email": req.Email, "password": req.Password})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", common.ErrEntityNotFound("User", err)
		}
		return "", common.ErrEntityNotFound("User", err)
	}

	token, _ := jwt.CreateToken(user.ID, user.Username)

	return token, nil
}

func (biz *userBiz) RegisterUser(ctx context.Context, req *request.RegisterUserRequest) error {
	var newUser model.UserCreate

	if err := copier.Copy(&newUser, req); err != nil {
		return err
	}
	err := biz.repo.Create(ctx, &newUser)
	return err
}

func (biz *userBiz) UpdateUser(ctx context.Context, id int, req *request.UpdateUserRequest) error {
	user, err := biz.repo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrEntityNotFound(model.User{}.EntityName(), err)
		}
		return common.ErrEntityNotFound(model.User{}.EntityName(), err)
	}

	if user.IsDeleted == 1 {
		return common.ErrEntityNotFound(model.User{}.EntityName(), nil)
	}

	var newUser model.UserUpdate
	if err := copier.Copy(&newUser, &req); err != nil {
		return err
	}

	err = biz.repo.Update(ctx, id, &newUser)
	return err
}

func (biz *userBiz) DeleteUser(ctx context.Context, id int) error {
	oldData, err := biz.repo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotDeleteEntity(model.User{}.EntityName(), err)
	}

	if oldData.IsDeleted == 1 {
		return common.ErrCannotDeleteEntity(model.User{}.EntityName(), err)
	}

	if err := biz.repo.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(model.User{}.EntityName(), err)
	}
	return nil
}

func (biz *userBiz) ListUser(ctx context.Context, filter *common.Filter, paging *common.Paging) ([]model.User, error) {
	result, err := biz.repo.ListWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.User{}.EntityName(), err)
	}

	return result, nil
}

func (biz *userBiz) FindUser(ctx context.Context, id int) (*model.User, error) {
	user, err := biz.repo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return nil, common.ErrEntityNotFound(model.User{}.EntityName(), err)
		}
		return nil, common.ErrEntityNotFound(model.User{}.EntityName(), err)
	}

	if user.IsDeleted == 1 {
		return nil, common.ErrEntityNotFound(model.User{}.EntityName(), err)
	}

	return user, nil
}
