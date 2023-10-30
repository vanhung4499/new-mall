package request

import (
	"new-mall/model"
	"new-mall/model/common"
)

type UserSearch struct {
	model.User
	common.PageInfo
}

type RegisterUserParam struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type UpdateUserInfoParam struct {
	NickName      string `json:"nickName"`
	PasswordMd5   string `json:"passwordMd5"`
	IntroduceSign string `json:"introduceSign"`
}

type UserLoginParam struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}
