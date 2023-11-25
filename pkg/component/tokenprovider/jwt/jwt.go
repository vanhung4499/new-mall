package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"new-mall/pkg/component/tokenprovider"
	"time"
)

type JwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *JwtProvider {
	return &JwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *JwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	// generate the JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now().UTC(),
	}, nil
}

func (j *JwtProvider) Validate(myToken string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}

func (j *JwtProvider) String() string {
	return "JWT implement Provider"
}
