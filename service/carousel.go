package service

import (
	"context"
	"new-mall/pkg/utils"
	"new-mall/repository/db/dao"
	"new-mall/types"
	"sync"
)

var CarouselSrvIns *CarouselSrv
var CarouselSrvOnce sync.Once

type CarouselSrv struct {
}

func GetCarouselSrv() *CarouselSrv {
	CarouselSrvOnce.Do(func() {
		CarouselSrvIns = &CarouselSrv{}
	})
	return CarouselSrvIns
}

func (s *CarouselSrv) ListCarousel(ctx context.Context, req *types.ListCarouselReq) (resp interface{}, err error) {
	carousels, err := dao.NewCarouselDao(ctx).ListCarousel()
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	resp = &types.DataListRes{
		Item:  carousels,
		Total: int64(len(carousels)),
	}

	return
}
