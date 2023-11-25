package types

import "new-mall/internal/models"

type UploadAvatarReq struct {
	Avatar string `form:"avatar" json:"avatar"`
}

type UserTokenRes struct {
	User         *models.User `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

type UserLoginReq struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type SendEmailServiceReq struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	// OperationType 1: Bind email 2: Unbind email 3: Change password
	OperationType uint `form:"operation_type" json:"operation_type"`
}

type ConfirmEmailReq struct {
	Token string `json:"token" form:"token"`
}
