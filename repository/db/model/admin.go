package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"new-mall/config"
	"new-mall/constant"
)

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}

func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	admin.PasswordDigest = string(bytes)
	return nil
}

func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordDigest), []byte(password))
	return err == nil
}

func (admin *Admin) AvatarURL() string {
	if config.Config.System.UploadModel == constant.UploadModelS3 {
		return admin.Avatar
	}
	local := config.Config.Local
	return local.Path + config.Config.System.HttpPort + local.StorePath + admin.Avatar
}
