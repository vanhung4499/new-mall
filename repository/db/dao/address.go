package dao

import (
	"context"
	"gorm.io/gorm"
	"new-mall/repository/db/model"
	"new-mall/types"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) GetAddressByAid(aId, uId uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).
		Where("id = ? AND user_id = ?", aId, uId).First(&address).
		Error

	return
}

func (dao *AddressDao) ListAddressByUid(uid uint) (r []*types.AddressRes, err error) {
	err = dao.DB.Model(&model.Address{}).
		Where("user_id = ?", uid).
		Order("created_at desc").
		Select("id, user_id, name, phone, address, UNIX_TIMESTAMP(created_at)").
		Find(&r).Error

	return
}

func (dao *AddressDao) CreateAddress(address *model.Address) (err error) {
	return dao.DB.Model(&model.Address{}).
		Create(&address).Error
}

func (dao *AddressDao) DeleteAddressById(aId, uId uint) (err error) {
	return dao.DB.Where("id = ? AND user_id = ?", aId, uId).
		Delete(&model.Address{}).Error
}

func (dao *AddressDao) UpdateAddressById(aId uint, address *model.Address) (err error) {
	return dao.DB.Model(&model.Address{}).
		Where("id=?", aId).
		Updates(address).Error
}
