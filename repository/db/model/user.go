package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"new-mall/config"
	"new-mall/constant"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          string
	Relations      []User `gorm:"many2many:relation;"`
}

const (
	PassWordCost        = 12       // Password encryption difficulty
	Active       string = "active" // Activate user
)

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

func (u *User) AvatarURL() string {
	if config.Config.System.UploadModel == constant.UploadModelS3 {
		return u.Avatar
	}
	local := config.Config.Local
	return local.Path + config.Config.System.HttpPort + local.StorePath + u.Avatar
}
