package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/internal/services"
	"new-mall/internal/types"
	"new-mall/pkg/common"
	"strconv"
)

type OrderController struct {
	OrderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{
		OrderService: orderService,
	}
}

func (c *OrderController) CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		err := c.OrderService.CreateOrderFromCart(ctx.Request.Context(), requester.GetID())
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *OrderController) ListOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging types.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		data, err := c.OrderService.ListOrder(ctx.Request.Context(), requester.GetID(), &paging)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, &paging))
	}
}

func (c *OrderController) GetOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		data, err := c.OrderService.GetOrder(ctx.Request.Context(), requester.GetID(), uint(id))
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *OrderController) DeleteOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		err = c.OrderService.DeleteOrder(ctx.Request.Context(), requester.GetID(), uint(id))
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
