package request

import (
	"new-mall/model"
	"new-mall/model/common"
)

type UserSearch struct {
	model.User
	common.PageInfo
}

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Avatar      string `json:"avatar"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
