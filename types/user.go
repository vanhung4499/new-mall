package types

type UserServiceReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // The front end makes judgments
}

type UserRegisterReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // The front end makes judgments
}

type UserTokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

type UserLoginReq struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type UserInfoUpdateReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
}

type UserInfoShowReq struct {
}

type UserFollowingReq struct {
	Id uint `json:"id" form:"id"`
}

type UserUnFollowingReq struct {
	Id uint `json:"id" form:"id"`
}

type SendEmailServiceReq struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	// OperationType 1: Bind email 2: Unbind email 3: Change password
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ValidEmailServiceReq struct {
	Token string `json:"token" form:"token"`
}

type UserInfoRes struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nickname"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}
