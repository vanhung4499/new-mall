package dao

import (
	"context"
	"gorm.io/gorm"
	"new-mall/pkg/utils"
	"new-mall/repository/db/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// FollowUser userId followed followerId
func (dao *UserDao) FollowUser(uId, followerId uint) (err error) {
	u, f := new(model.User), new(model.User)
	dao.DB.Model(&model.User{}).Where(`id = ?`, uId).First(&u)
	dao.DB.Model(&model.User{}).Where(`id = ?`, followerId).First(&f)
	err = dao.DB.Model(&f).Association(`Relations`).
		Append([]model.User{*u})
	if err != nil {
		utils.Logger.Error(err)
		return err
	}

	return
}

func (dao *UserDao) UnFollowUser(uId, followerId uint) (err error) {
	u, f := new(model.User), new(model.User)
	dao.DB.Model(&model.User{}).Where(`id = ?`, uId).First(&u)
	dao.DB.Model(&model.User{}).Where(`id = ?`, followerId).First(&f)
	err = dao.DB.Model(&u).Association(`Relations`).Delete(f)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	return
}

func (dao *UserDao) ListFollowing(userId uint) (f []*model.User, err error) {
	u := new(model.User)
	f = make([]*model.User, 0)
	dao.DB.Model(&model.User{}).Where(`id = ?`, userId).First(&u)
	err = dao.DB.Model(&u).Association(`Relations`).
		Find(&f)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	return
}

func (dao *UserDao) ListFollower(userId int64) (f []*model.User, err error) {
	u := new(model.User)
	f = make([]*model.User, 0)
	dao.DB.Model(&model.User{}).Where(`id = ?`, userId).First(&u)
	err = dao.DB.Model(&f).Association(`Relations`).
		Find(&u)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	return
}
func (dao *UserDao) GetUserById(uId uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}

func (dao *UserDao) UpdateUserById(uId uint, user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}
