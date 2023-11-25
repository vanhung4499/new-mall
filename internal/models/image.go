package models

import (
	"errors"
	"gorm.io/gorm"
	"new-mall/pkg/common"
)

const ImageEntityName = "Image"

type Image struct {
	gorm.Model
	ProductID uint `gorm:"not null"`
	ImageURL  string
}

func (Image) TableName() string {
	return "images"
}

type ImageCreate struct {
	ProductID uint   `form:"product_id" binding:"required"`
	ImageURL  string `form:"image_url" binding:"required"`
}

func (ImageCreate) TableName() string {
	return Image{}.TableName()
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge",
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}
