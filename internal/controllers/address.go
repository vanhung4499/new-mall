package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/internal/models"
	"new-mall/internal/services"
	"new-mall/pkg/common"
	"strconv"
)

type AddressController struct {
	AddressService *services.AddressService
}

func NewAddressController(addressService *services.AddressService) *AddressController {
	return &AddressController{
		AddressService: addressService,
	}
}

func (c *AddressController) CreateAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.AddressCreate
		if err := ctx.ShouldBind(&data); err != nil {
			// Parameter verification
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.AddressService.CreateAddress(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *AddressController) GetAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data, err := c.AddressService.GetAddress(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *AddressController) ListAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		conditions := map[string]interface{}{"user_id": requester.GetID()}

		data, err := c.AddressService.ListAddress(ctx.Request.Context(), conditions, nil, nil)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *AddressController) UpdateAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data models.AddressUpdate
		if err = ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		data.UserID = requester.GetID()

		if err = c.AddressService.UpdateAddress(ctx.Request.Context(), uint(id), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *AddressController) DeleteAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		if err = c.AddressService.DeleteAddress(ctx.Request.Context(), uint(id), requester.GetID()); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
