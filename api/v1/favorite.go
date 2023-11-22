package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/constant"
	"new-mall/pkg/utils"
	"new-mall/pkg/utils/ctl"
	"new-mall/service"
	"new-mall/types"
)

func CreateFavoriteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoriteCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetFavoriteSrv()
		res, err := l.FavoriteCreate(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func ListFavoritesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoritesServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		if req.PageSize == 0 {
			req.PageSize = constant.BasePageSize
		}

		l := service.GetFavoriteSrv()
		res, err := l.FavoriteList(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func DeleteFavoriteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.FavoriteDeleteReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetFavoriteSrv()
		res, err := l.FavoriteDelete(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}
