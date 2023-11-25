package middleware

import (
	"github.com/gin-gonic/gin"
	"new-mall/internal/global"
	"new-mall/internal/repositories"
	"new-mall/pkg/common"
	"new-mall/pkg/component/tokenprovider/jwt"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"wrong authentication header",
		"ErrWrongAuthHeader",
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth() func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(global.CONFIG.EncryptSecret.JwtSecret)

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		db := global.DB
		repo := repositories.NewUserRepository(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := repo.FindWithCondition(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
