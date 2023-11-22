package service

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"new-mall/pkg/utils"
	"new-mall/pkg/utils/ctl"
	"new-mall/repository/cache"
	"new-mall/repository/db/dao"
	"new-mall/repository/db/model"
	"new-mall/types"
	"strconv"
	"sync"
	"time"
)

const OrderTimeKey = "OrderTime"

var OrderSrvIns *OrderSrv
var OrderSrvOnce sync.Once

type OrderSrv struct {
}

func GetOrderSrv() *OrderSrv {
	OrderSrvOnce.Do(func() {
		OrderSrvIns = &OrderSrv{}
	})
	return OrderSrvIns
}

func (s *OrderSrv) OrderCreate(ctx context.Context, req *types.OrderCreateReq) (res interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	order := &model.Order{
		UserID:    u.Id,
		ProductID: req.ProductID,
		BossID:    req.BossID,
		Num:       int(req.Num),
		Money:     float64(req.Money),
		Type:      1,
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(req.AddressID, u.Id)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	order.AddressID = address.ID
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(req.ProductID))
	userNum := strconv.Itoa(int(u.Id))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum

	orderDao := dao.NewOrderDao(ctx)
	err = orderDao.CreateOrder(order)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	// Store the order number in Redis and set the expiration time
	data := redis.Z{
		Score:  float64(time.Now().Unix()) + 15*time.Minute.Seconds(),
		Member: orderNum,
	}
	cache.RedisClient.ZAdd(cache.RedisContext, OrderTimeKey, data)

	return
}

func (s *OrderSrv) OrderList(ctx context.Context, req *types.OrderListReq) (res interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	orders, total, err := dao.NewOrderDao(ctx).ListOrderByCondition(u.Id, req)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	res = types.DataListRes{
		Item:  orders,
		Total: total,
	}

	return
}

func (s *OrderSrv) OrderShow(ctx context.Context, req *types.OrderShowReq) (res interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	order, err := dao.NewOrderDao(ctx).ShowOrderById(req.OrderId, u.Id)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	res = order

	return
}

func (s *OrderSrv) OrderDelete(ctx context.Context, req *types.OrderDeleteReq) (res interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	err = dao.NewOrderDao(ctx).DeleteOrderById(req.OrderId, u.Id)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	return
}
