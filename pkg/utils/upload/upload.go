package upload

import (
	"mime/multipart"
	"new-mall/config"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewOss() OSS {
	switch config.Config.System.UploadModel {
	case "local":
		return &Local{}
	case "aws-s3":
		return &AwsS3{}
	default:
		return &Local{}
	}
}
