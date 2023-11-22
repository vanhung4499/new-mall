package middleware

import (
	"github.com/gin-gonic/gin"
	"new-mall/constant"
	"new-mall/pkg/e"
	"new-mall/pkg/utils"
	"new-mall/pkg/utils/ctl"
)

// AuthMiddleware is the token authentication middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		accessToken := c.GetHeader("access_token")
		refreshToken := c.GetHeader("refresh_token")
		if accessToken == "" {
			code = e.InvalidParams
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token cannot be empty",
			})
			c.Abort()
			return
		}
		newAccessToken, newRefreshToken, err := utils.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Authentication failed",
				"error":  err.Error(),
			})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(newAccessToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   err.Error(),
			})
			c.Abort()
			return
		}
		SetToken(c, newAccessToken, newRefreshToken)
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.ID}))
		ctl.InitUserInfo(c.Request.Context())
		c.Next()
	}
}

func SetToken(c *gin.Context, accessToken, refreshToken string) {
	secure := IsHttps(c)
	c.Header(constant.AccessTokenHeader, accessToken)
	c.Header(constant.RefreshTokenHeader, refreshToken)
	c.SetCookie(constant.AccessTokenHeader, accessToken, constant.MaxAge, "/", "", secure, true)
	c.SetCookie(constant.RefreshTokenHeader, refreshToken, constant.MaxAge, "/", "", secure, true)
}

// IsHttps checks if the request is using HTTPS
func IsHttps(c *gin.Context) bool {
	if c.GetHeader(constant.HeaderForwardedProto) == "https" || c.Request.TLS != nil {
		return true
	}
	return false
}
