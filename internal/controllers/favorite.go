package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/internal/models"
	"new-mall/internal/services"
	"new-mall/internal/types"
	"new-mall/pkg/common"
	"strconv"
)

type FavoriteController struct {
	FavoriteService *services.FavoriteService
}

func NewFavoriteController(favoriteService *services.FavoriteService) *FavoriteController {
	return &FavoriteController{
		FavoriteService: favoriteService,
	}
}

func (c *FavoriteController) CreateFavorite() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.FavoriteCreate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.FavoriteService.CreateFavorite(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *FavoriteController) GetFavorite() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data, err := c.FavoriteService.GetFavorite(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *FavoriteController) ListFavorite() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging types.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		condition := map[string]interface{}{"user_id": requester.GetID()}

		data, err := c.FavoriteService.ListFavorite(ctx.Request.Context(), condition, nil, &paging)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, paging))
	}
}

func (c *FavoriteController) DeleteFavorite() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err = c.FavoriteService.DeleteFavorite(ctx.Request.Context(), uint(id)); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
