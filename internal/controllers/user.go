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

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) RegisterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err := c.UserService.RegisterUser(ctx.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *UserController) LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserLoginReq
		if err := ctx.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		res, err := c.UserService.LoginUser(ctx.Request.Context(), &req)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}

func (c *UserController) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.UserUpdate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		user := ctx.MustGet(common.CurrentUser).(*models.User)

		err := c.UserService.UpdateUser(ctx.Request.Context(), user.ID, &data)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (c *UserController) GetProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		data, err := c.UserService.GetUser(ctx.Request.Context(), requester.GetID())
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *UserController) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data, err := c.UserService.GetUser(ctx.Request.Context(), uint(id))
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func (c *UserController) UploadAvatar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UploadAvatarReq
		if err := ctx.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		_, fileHeader, err := ctx.Request.FormFile("file")
		if fileHeader == nil {
			panic(models.ErrFileIsNotImage(err))
		}

		user := ctx.MustGet(common.CurrentUser).(*models.User)

		err = c.UserService.UploadUserAvatar(ctx.Request.Context(), user.ID, fileHeader)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
