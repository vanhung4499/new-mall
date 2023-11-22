package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"new-mall/config"
	"new-mall/constant"
	"time"
)

var jwtSecret = []byte(config.Config.EncryptSecret.JwtSecret)

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken generates user access and refresh tokens
func GenerateToken(id uint, username string) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(constant.AccessTokenExpireDuration)
	rtExpireTime := nowTime.Add(constant.RefreshTokenExpireDuration)

	// Create claims for the access token
	claims := Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mall",
		},
	}

	// Generate and sign the access token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Generate and sign the refresh token
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: rtExpireTime.Unix(),
		Issuer:    "mall",
	}).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// ParseToken validates and parses the user access token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// ParseRefreshToken validates and refreshes user tokens
func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// Parse the access token
	accessClaim, err := ParseToken(aToken)
	if err != nil {
		return
	}

	// Parse the refresh token
	refreshClaim, err := ParseToken(rToken)
	if err != nil {
		return
	}

	// Check if the access token is not expired
	if accessClaim.ExpiresAt > time.Now().Unix() {
		// If the access token is not expired, refresh both access and refresh tokens
		return GenerateToken(accessClaim.ID, accessClaim.Username)
	}

	// Check if the refresh token is not expired
	if refreshClaim.ExpiresAt > time.Now().Unix() {
		// If the access token is expired but the refresh token is not, refresh both access and refresh tokens
		return GenerateToken(accessClaim.ID, accessClaim.Username)
	}

	// If both tokens are expired, return an error indicating the need to re-login
	return "", "", errors.New("identity expired, please log in again")
}

// EmailClaims represents claims for email verification tokens
type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// GenerateEmailToken generates email verification tokens
func GenerateEmailToken(userID, Operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * time.Minute)
	claims := EmailClaims{
		UserID:        userID,
		Email:         email,
		Password:      password,
		OperationType: Operation,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mall",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseEmailToken validates and parses email verification tokens
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
