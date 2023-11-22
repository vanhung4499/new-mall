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

func CreateAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetAddressSrv()
		res, err := l.AddressCreate(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func ShowAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressGetReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetAddressSrv()
		res, err := l.AddressShow(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func ListAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressListReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		if req.PageSize == 0 {
			req.PageSize = constant.BasePageSize
		}

		l := service.GetAddressSrv()
		res, err := l.AddressList(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func UpdateAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetAddressSrv()
		res, err := l.AddressUpdate(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func DeleteAddressHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddressDeleteReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetAddressSrv()
		res, err := l.AddressDelete(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}
