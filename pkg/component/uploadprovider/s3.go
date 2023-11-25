package uploadprovider

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"new-mall/internal/global"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsS3 struct{}

// UploadFile uploads a file to AWS S3
func (*AwsS3) UploadFile(ctx context.Context, file *multipart.FileHeader) (string, string, error) {
	session := newSession()
	uploader := s3manager.NewUploader(session)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := global.CONFIG.AwsS3.PathPrefix + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close()

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(global.CONFIG.AwsS3.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		return "", "", err
	}

	return global.CONFIG.AwsS3.BaseURL + "/" + filename, fileKey, nil
}

// DeleteFile deletes a file from AWS S3
func (*AwsS3) DeleteFile(ctx context.Context, key string) error {
	session := newSession()
	svc := s3.New(session)
	filename := global.CONFIG.AwsS3.PathPrefix + "/" + key
	bucket := global.CONFIG.AwsS3.Bucket

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return errors.New("function svc.DeleteObject() failed, err:" + err.Error())
	}

	_ = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return nil
}

// newSession creates an AWS S3 session
func newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String(global.CONFIG.AwsS3.Region),
		Endpoint:         aws.String(global.CONFIG.AwsS3.Endpoint),
		S3ForcePathStyle: aws.Bool(global.CONFIG.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(global.CONFIG.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			global.CONFIG.AwsS3.SecretID,
			global.CONFIG.AwsS3.SecretKey,
			"",
		),
	})
	return sess
}
