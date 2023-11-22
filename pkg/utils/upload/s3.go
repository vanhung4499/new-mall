package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"new-mall/config"
	"new-mall/pkg/utils"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsS3 struct{}

// UploadFile uploads a file to AWS S3
func (*AwsS3) UploadFile(file *multipart.FileHeader) (string, string, error) {
	session := newSession()
	uploader := s3manager.NewUploader(session)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := config.Config.AwsS3.PathPrefix + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		utils.Logger.Error("function file.Open() failed", openError.Error())
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close()

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Config.AwsS3.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		utils.Logger.Error("function uploader.Upload() failed", err.Error())
		return "", "", err
	}

	return config.Config.AwsS3.BaseURL + "/" + filename, fileKey, nil
}

// DeleteFile deletes a file from AWS S3
func (*AwsS3) DeleteFile(key string) error {
	session := newSession()
	svc := s3.New(session)
	filename := config.Config.AwsS3.PathPrefix + "/" + key
	bucket := config.Config.AwsS3.Bucket

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		utils.Logger.Error("function svc.DeleteObject() failed", err.Error())
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
		Region:           aws.String(config.Config.AwsS3.Region),
		Endpoint:         aws.String(config.Config.AwsS3.Endpoint),
		S3ForcePathStyle: aws.Bool(config.Config.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(config.Config.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			config.Config.AwsS3.SecretID,
			config.Config.AwsS3.SecretKey,
			"",
		),
	})
	return sess
}
