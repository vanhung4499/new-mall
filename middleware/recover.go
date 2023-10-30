package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"new-mall/model/common"
)

func Recover(c context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.Error); ok {
					c.AbortWithStatusJSON(appErr.Code, appErr)
					panic(err)
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.Code, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
