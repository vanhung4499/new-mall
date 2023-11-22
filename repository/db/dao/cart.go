package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"new-mall/pkg/e"
	"new-mall/repository/db/model"
	"new-mall/types"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

// CreateCart creates cart pId (product id), uId (user id), bId (store id)
func (dao *CartDao) CreateCart(pId, uId, bId uint) (cart *model.Cart, status int, err error) {
	// Check if this product is available
	cart, err = dao.GetCartById(pId, uId, bId)
	// Empty, joining for the first time
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cart = &model.Cart{
			UserID:    uId,
			ProductID: pId,
			BossID:    bId,
			Num:       1,
			MaxNum:    10,
			Check:     false,
		}
		err = dao.DB.Create(&cart).Error
		if err != nil {
			return
		}
		return cart, e.SUCCESS, err
	}
	if cart.Num < cart.MaxNum {
		// less than maximum num
		cart.Num++
		err = dao.DB.Save(&cart).Error
		if err != nil {
			return
		}
		return cart, e.ErrorProductExistCart, err
	}
	// greater than the maximum num
	return cart, e.ErrorProductMoreCart, err
}

// GetCartById Get Cart by ID
func (dao *CartDao) GetCartById(pId, uId, bId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).
		Where("user_id = ? AND product_id = ? AND boss_id = ?",
			uId, pId, bId).
		First(&cart).Error

	return
}

// ListCartByUserId gets Cart by user_id
func (dao *CartDao) ListCartByUserId(uId uint) (cart []*types.CartRes, err error) {
	err = dao.DB.Model(&model.Cart{}).
		Joins("AS c LEFT JOIN product AS p ON c.product_id = p.id").
		Where("c.user_id = ?", uId).
		Select("c.id AS id," +
			"c.user_id AS user_id," +
			"c.product_id AS product_id," +
			"UNIX_TIMESTAMP(c.created_at) AS created_at," +
			"c.num AS num," +
			"c.max_num AS max_num," +
			"c.check AS check_," +
			"p.img_path AS img_path," +
			"p.boss_id AS boss_id," +
			"p.boss_name AS boss_name," +
			"p.info AS info," +
			"p.discount_price AS discount_price").
		Find(&cart).Error

	return
}

func (dao *CartDao) UpdateCartNumById(cId, uId, num uint) error {
	return dao.DB.Model(&model.Cart{}).
		Where("id = ? AND user_id = ?", cId, uId).
		Update("num", num).Error
}

func (dao *CartDao) DeleteCartById(cId, uId uint) error {
	return dao.DB.Model(&model.Cart{}).
		Where("id = ? AND user_id = ?", cId, uId).
		Delete(&model.Cart{}).Error
}
