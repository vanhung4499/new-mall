package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/internal/models"
	"new-mall/internal/services"
	"new-mall/pkg/common"
	"strconv"
)

type CategoryController struct {
	CategoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (c *CategoryController) CreateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.CategoryCreate
		if err := ctx.ShouldBind(&data); err != nil {
			// Parameter verification
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.CategoryService.CreateCategory(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *CategoryController) GetCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data, err := c.CategoryService.GetCategory(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *CategoryController) ListCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := c.CategoryService.ListCategory(ctx.Request.Context(), nil, nil)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}

func (c *CategoryController) UpdateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data models.CategoryUpdate
		if err = ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err = c.CategoryService.UpdateCategory(ctx.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *CategoryController) DeleteCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err = c.CategoryService.DeleteCategory(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
