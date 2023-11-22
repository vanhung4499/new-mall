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

func CreateOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetOrderSrv()
		resp, err := l.OrderCreate(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, resp))
		return
	}
}

func ListOrdersHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderListReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		if req.PageSize == 0 {
			req.PageSize = constant.BasePageSize
		}

		l := service.GetOrderSrv()
		resp, err := l.OrderList(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, resp))
		return
	}
}

// ShowOrderHandler 订单详情
func ShowOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderShowReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetOrderSrv()
		resp, err := l.OrderShow(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, resp))
		return
	}
}

func DeleteOrderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.OrderDeleteReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetOrderSrv()
		resp, err := l.OrderDelete(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, resp))
	}
}
