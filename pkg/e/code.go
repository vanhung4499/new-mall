package e

const (
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	NotExistIdentifier    = 202
	ERROR                 = 500
	InvalidParams         = 400

	// User Error
	ErrorExistNick          = 10001
	ErrorExistUser          = 10002
	ErrorNotExistUser       = 10003
	ErrorNotCompare         = 10004
	ErrorNotComparePassword = 10005
	ErrorFailEncryption     = 10006
	ErrorNotExistProduct    = 10007
	ErrorNotExistAddress    = 10008
	ErrorExistFavorite      = 10009
	ErrorUserNotFound       = 10010

	// Store Error
	ErrorBossCheckTokenFail        = 20001
	ErrorBossCheckTokenTimeout     = 20002
	ErrorBossToken                 = 20003
	ErrorBoss                      = 20004
	ErrorBossInsufficientAuthority = 20005
	ErrorBossProduct               = 20006

	// Cart Error
	ErrorProductExistCart = 20007
	ErrorProductMoreCart  = 20008

	//Administrator Error
	ErrorAuthCheckTokenFail        = 30001 //token fail
	ErrorAuthCheckTokenTimeout     = 30002 //token Expired
	ErrorAuthToken                 = 30003
	ErrorAuth                      = 30004
	ErrorAuthInsufficientAuthority = 30005
	ErrorReadFile                  = 30006
	ErrorSendEmail                 = 30007
	ErrorCallApi                   = 30008
	ErrorUnmarshalJson             = 30009
	ErrorAdminFindUser             = 30010

	// Database Error
	ErrorDatabase = 40001

	// Storage Error
	ErrorAwsS3      = 50001
	ErrorUploadFile = 50002
)
