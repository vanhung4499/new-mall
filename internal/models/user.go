package models

import (
	"errors"
	"gorm.io/gorm"
	"new-mall/pkg/common"
)

const UserEntityName = "User"

type User struct {
	gorm.Model
	Email     string `gorm:"unique,not null"`
	Username  string `gorm:"unique"`
	Password  string `json:"-" gorm:"not null"`
	Salt      string `json:"-" gorm:"not null"`
	FirstName string
	LastName  string
	Phone     string `gorm:"unique"`
	Status    string `gorm:"default:active"`
	Avatar    string
	Role      common.AppRole
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role.String()
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	gorm.Model
	Username  string         `form:"username,omitempty"`
	Email     string         `form:"email" binding:"required"`
	Password  string         `form:"password" binding:"required"`
	Salt      string         `json:"-"`
	FirstName string         `form:"first_name,omitempty"`
	LastName  string         `form:"last_name,omitempty"`
	Phone     string         `form:"phone,omitempty"`
	Status    string         `form:"status,omitempty"`
	Avatar    string         `form:"avatar"`
	Role      common.AppRole `form:"role,omitempty"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserUpdate struct {
	Username  string `form:"username,omitempty"`
	FirstName string `form:"first_name,omitempty"`
	LastName  string `form:"last_name,omitempty"`
	Phone     string `form:"phone,omitempty"`
	Status    string `form:"status,omitempty"`
	Avatar    string `form:"avatar,omitempty"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrUsernameExisted = common.NewCustomError(
		errors.New("username has already existed"),
		"username has already existed",
		"ErrUsernameExisted",
	)
)
