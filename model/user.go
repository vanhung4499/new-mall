package model

import (
	"new-mall/global"
	"new-mall/utils"
)

type User struct {
	global.Model
	Email       string `json:"email" form:"email" gorm:"column:email;unique;type:varchar(50);"`
	Username    string `json:"username" form:"username" gorm:"column:username;unique;type:varchar(50);"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" gorm:"column:phone_number;type:varchar(20);"`
	Password    string `json:"password" form:"password" gorm:"column:password;type:varchar(32);"`
	Avatar      string `json:"avatar" form:"avatar" gorm:"column:avatar;type:varchar(1000);"`
	LockedFlag  int    `json:"lockedFlag" form:"lockedFlag" gorm:"column:locked_flag;type:tinyint"`
}

// TableName User
func (User) TableName() string {
	return "users"
}

// SetPassword Set Password
func (u *User) SetPassword(password string) {
	u.Password = utils.MD5V([]byte(password))
}

type UserCreate struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type UserUpdate struct {
	NickName      string `json:"nickName"`
	PasswordMd5   string `json:"passwordMd5"`
	IntroduceSign string `json:"introduceSign"`
}
