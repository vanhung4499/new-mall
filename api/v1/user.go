package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/constant"
	"new-mall/pkg/e"
	"new-mall/pkg/utils"
	"new-mall/pkg/utils/ctl"
	"new-mall/service"
	"new-mall/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		// Parameter verification
		if req.Key == "" || len(req.Key) != constant.EncryptMoneyKeyLength {
			err := errors.New("key length is wrong, must be 6 digits")
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		res, err := l.UserRegister(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserLoginReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		res, err := l.UserLogin(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func UserUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoUpdateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		res, err := l.UserInfoUpdate(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func ShowUserInfoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoShowReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		res, err := l.UserInfoShow(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func UploadAvatarHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		_, fileHeader, _ := ctx.Request.FormFile("file")
		if fileHeader == nil {
			err := errors.New(e.GetMsg(e.ErrorUploadFile))
			ctx.JSON(constant.IllegalRequest, ErrorResponse(ctx, err))
			utils.Logger.Infoln(err)
			return
		}

		l := service.GetUserSrv()
		res, err := l.UserAvatarUpload(ctx.Request.Context(), fileHeader, &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func SendEmailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SendEmailServiceReq

		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		res, err := l.SendEmail(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func UserFollowingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserFollowingReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		res, err := l.UserFollow(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func UserUnFollowingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserUnFollowingReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		res, err := l.UserUnFollow(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}

func ValidEmailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ValidEmailServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			// Parameter verification
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetUserSrv()
		res, err := l.Valid(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.ResSuccess(ctx, res))
	}
}
