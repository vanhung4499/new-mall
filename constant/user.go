package constant

import "time"

const (
	EmailOperationBinding = iota + 1
	EmailOperationNoBinding
	EmailOperationUpdatePassword
)

var EmailOperationMap = map[uint]string{
	EmailOperationBinding:        "You are binding your email address, please click the link to confirm your identity %s",
	EmailOperationNoBinding:      "You are in mailbox, please click the link to confirm your identity %s",
	EmailOperationUpdatePassword: "You are changing your password, please click the link to verify your identity %s",
}

const (
	AccessTokenHeader    = "access_token"
	RefreshTokenHeader   = "refresh_token"
	HeaderForwardedProto = "X-Forwarded-Proto"
	MaxAge               = 3600 * 24
)

const (
	AccessToken = iota
	RefreshToken
)

const (
	AccessTokenExpireDuration  = 24 * time.Hour
	RefreshTokenExpireDuration = 10 * 24 * time.Hour
)

const EncryptMoneyKeyLength = 6

const UserInitMoney = "10000" // Initial amount 10K

const (
	UserDefaultAvatarOss   = "https://api.dicebear.com/avatar.svg" // OSS’s default avatar
	UserDefaultAvatarLocal = "avatar.JPG"                          // Local default avatar

)
