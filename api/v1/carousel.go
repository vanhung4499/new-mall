package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/pkg/utils"
	"new-mall/pkg/utils/ctl"
	"new-mall/service"
	"new-mall/types"
)

func ListCarouselsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListCarouselReq
		if err := ctx.ShouldBind(&req); err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetCarouselSrv()
		resp, err := l.ListCarousel(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, resp))
	}
}
