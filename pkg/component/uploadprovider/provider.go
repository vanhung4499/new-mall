package uploadprovider

import (
	"context"
	"mime/multipart"
	"new-mall/internal/global"
)

type UploadProvider interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) (string, string, error)
	DeleteFile(ctx context.Context, key string) error
}

func NewUploadProvider() UploadProvider {
	switch global.CONFIG.System.UploadModel {
	case "local":
		return &Local{}
	case "aws-s3":
		return &AwsS3{}
	default:
		return &Local{}
	}
}
