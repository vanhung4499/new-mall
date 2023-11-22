package dao

import (
	"context"
	"gorm.io/gorm"
	"new-mall/repository/db/model"
	"new-mall/types"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewNewCarouselDao(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

func (dao *CarouselDao) ListCarousel() (r []*types.ListCarouselResp, err error) {
	err = dao.DB.Model(&model.Carousel{}).
		Select("id, image_path, product_id, UNIX_TIMESTAMP(created_at)").
		Find(&r).Error

	return
}
