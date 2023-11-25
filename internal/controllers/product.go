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

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (c *ProductController) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.ProductCreate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ProductService.CreateProduct(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.ID))
	}
}

func (c *ProductController) ListProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter types.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var paging types.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		result, err := c.ProductService.ListProduct(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

func (c *ProductController) GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data, err := c.ProductService.GetProduct(ctx.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *ProductController) DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err = c.ProductService.DeleteProduct(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *ProductController) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data models.ProductUpdate
		if err = ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err = c.ProductService.UpdateProduct(ctx.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
