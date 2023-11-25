package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/internal/models"
	"new-mall/internal/services"
	"new-mall/pkg/common"
	"strconv"
)

type CarouselController struct {
	CarouselService *services.CarouselService
}

func NewCarouselController(carouselService *services.CarouselService) *CarouselController {
	return &CarouselController{
		CarouselService: carouselService,
	}
}

func (c *CarouselController) CreateCarousel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.CarouselCreate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// TODO: Check if the user has permission to create this carousel

		if err := c.CarouselService.CreateCarousel(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}

func (c *CarouselController) GetCarousel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data, err := c.CarouselService.GetCarousel(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *CarouselController) ListCarousel() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		data, err := c.CarouselService.ListCarousel(ctx.Request.Context(), nil, nil)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *CarouselController) UpdateCarousel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// TODO: Check if the user has permission to update this carousel

		var data models.CarouselUpdate
		if err = ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err = c.CarouselService.UpdateCarousel(ctx.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *CarouselController) DeleteCarousel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// TODO: Check if the user has permission to delete this carousel

		if err = c.CarouselService.DeleteCarousel(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
