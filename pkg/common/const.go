package common

const (
	CurrentUser = "currentUser"

	DateTimeFormat = "2006-01-02 15:04:05.999999 -07:00"
)

const (
	OrderTypeUnPaid = iota + 1
	OrderTypePendingShipping
	OrderTypeShipping
	OrderTypeReceipt
)

var OrderTypeMap = map[int]string{
	OrderTypeUnPaid:          "Unpaid",
	OrderTypePendingShipping: "Paid, waiting for shipment",
	OrderTypeShipping:        "Shipped, waiting for receipt",
	OrderTypeReceipt:         "Goods received, transaction successful",
}

const (
	UploadModelS3    = "s3"
	UploadModelLocal = "local"
)

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
	UserDefaultAvatarOss   = "https://api.dicebear.com/avatar.svg"          // OSS’s default avatar
	UserDefaultAvatarLocal = "http:localhost:5001/static/upload/avatar.JPG" // Local default avatar
)
