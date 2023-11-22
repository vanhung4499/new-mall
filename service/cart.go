package service

import (
	"context"
	"errors"
	"new-mall/pkg/e"
	"new-mall/pkg/utils"
	"new-mall/pkg/utils/ctl"
	"new-mall/repository/db/dao"
	"new-mall/types"
	"sync"
)

var CartSrvIns *CartSrv
var CartSrvOnce sync.Once

type CartSrv struct {
}

func GetCartSrv() *CartSrv {
	CartSrvOnce.Do(func() {
		CartSrvIns = &CartSrv{}
	})
	return CartSrvIns
}

func (s *CartSrv) CartCreate(ctx context.Context, req *types.CartCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	// Determine whether this product is available
	_, err = dao.NewProductDao(ctx).GetProductById(req.ProductId)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	// Determine whether this product is available
	cartDao := dao.NewCartDao(ctx)
	_, status, _ := cartDao.CreateCart(req.ProductId, u.Id, req.BossID)
	if status == e.ErrorProductMoreCart {
		err = errors.New(e.GetMsg(status))
		return
	}
	return
}

func (s *CartSrv) CartList(ctx context.Context, req *types.CartListReq) (res interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	carts, err := dao.NewCartDao(ctx).ListCartByUserId(u.Id)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	res = &types.DataListRes{
		Item:  carts, // TODO no paging, consider whether to add it later
		Total: int64(len(carts)),
	}

	return
}

func (s *CartSrv) CartUpdate(ctx context.Context, req *types.UpdateCartServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	err = dao.NewCartDao(ctx).UpdateCartNumById(req.Id, u.Id, req.Num)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	return
}

func (s *CartSrv) CartDelete(ctx context.Context, req *types.CartDeleteReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	err = dao.NewCartDao(ctx).DeleteCartById(req.Id, u.Id)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	return
}
