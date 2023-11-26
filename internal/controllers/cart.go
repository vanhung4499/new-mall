package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/internal/models"
	"new-mall/internal/services"
	"new-mall/pkg/common"
	"strconv"
)

type CartController struct {
	CartService *services.CartService
}

func NewCartController(cartService *services.CartService) *CartController {
	return &CartController{
		CartService: cartService,
	}
}

func (c *CartController) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var data models.CartItemCreate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		err := c.CartService.AddToCart(ctx.Request.Context(), requester.GetID(), &data)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *CartController) GetCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		cart, err := c.CartService.GetCart(ctx.Request.Context(), requester.GetID())
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(cart))
	}
}

func (c *CartController) UpdateCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.CartItemUpdate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		err := c.CartService.UpdateCart(ctx.Request.Context(), requester.GetID(), &data)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *CartController) DeleteCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		err := c.CartService.DeleteCart(ctx.Request.Context(), requester.GetID())
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *CartController) DeleteCartItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productId, err := strconv.Atoi(ctx.Param("productId"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		err = c.CartService.DeleteCartItemWithCondition(ctx.Request.Context(), requester.GetID(), uint(productId))
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
