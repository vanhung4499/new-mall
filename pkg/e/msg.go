package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	UpdatePasswordSuccess: "Password has been updated",
	NotExistIdentifier:    "The third-party account is not bound",
	ERROR:                 "fail",
	InvalidParams:         "Request parameter error",

	ErrorExistNick:          "This nickname already exists",
	ErrorExistUser:          "This username already exists",
	ErrorNotExistUser:       "This user does not exist",
	ErrorNotCompare:         "Account password is wrong",
	ErrorNotComparePassword: "The two password inputs do not match",
	ErrorFailEncryption:     "Encryption failed",
	ErrorNotExistProduct:    "This product does not exist",
	ErrorNotExistAddress:    "The shipping address does not exist",
	ErrorExistFavorite:      "The product has already been favorite",
	ErrorUserNotFound:       "User not found",

	ErrorBossCheckTokenFail:        "Boss's token authentication failed",
	ErrorBossCheckTokenTimeout:     "Boss's token has timed out",
	ErrorBossToken:                 "Failed to generate Boss's token",
	ErrorBoss:                      "Boss's token error",
	ErrorBossInsufficientAuthority: "Insufficient authority for Boss",
	ErrorBossProduct:               "Boss encountered a file reading error",

	ErrorProductExistCart: "The product is already in the shopping cart, quantity +1",
	ErrorProductMoreCart:  "Exceeded the maximum limit",

	ErrorAuthCheckTokenFail:        "Token authentication failed",
	ErrorAuthCheckTokenTimeout:     "Token has timed out",
	ErrorAuthToken:                 "Failed to generate token",
	ErrorAuth:                      "Token error",
	ErrorAuthInsufficientAuthority: "Insufficient authority",
	ErrorReadFile:                  "Failed to read the file",
	ErrorSendEmail:                 "Failed to send email",
	ErrorCallApi:                   "Failed to call the API",
	ErrorUnmarshalJson:             "Failed to decode JSON",

	ErrorUploadFile:    "Upload failed",
	ErrorAdminFindUser: "Admin failed to find user",

	ErrorDatabase: "Database operation error, please try again",

	ErrorAwsS3: "S3 configuration error",
}

// GetMsg retrieves the corresponding message for a given status code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
