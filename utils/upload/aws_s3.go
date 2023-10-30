package upload

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
	"mime/multipart"
	"new-mall/global"
	"time"
)

type AwsS3 struct{}

//@author: [WqyJh](https://github.com/WqyJh)
//@object: *AwsS3
//@function: UploadFile
//@description: Upload file to Aws S3 using aws-sdk-go. See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html#s3-examples-bucket-ops-upload-file-to-bucket
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*AwsS3) UploadFile(file *multipart.FileHeader) (string, string, error) {
	session := newSession()
	uploader := s3manager.NewUploader(session)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := global.Config.AwsS3.PathPrefix + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		global.Log.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // Create file defer close

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(global.Config.AwsS3.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		global.Log.Error("function uploader.Upload() failed", zap.Any("err", err.Error()))
		return "", "", err
	}

	return global.Config.AwsS3.BaseURL + "/" + filename, fileKey, nil
}

//@author: [WqyJh](https://github.com/WqyJh)
//@object: *AwsS3
//@function: DeleteFile
//@description: Delete file from Aws S3 using aws-sdk-go. See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html#s3-examples-bucket-ops-delete-bucket-item
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*AwsS3) DeleteFile(key string) error {
	session := newSession()
	svc := s3.New(session)
	filename := global.Config.AwsS3.PathPrefix + "/" + key
	bucket := global.Config.AwsS3.Bucket

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		global.Log.Error("function svc.DeleteObject() failed", zap.Any("err", err.Error()))
		return errors.New("function svc.DeleteObject() failed, err:" + err.Error())
	}

	_ = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return nil
}

// newSession Create S3 session
func newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String(global.Config.AwsS3.Region),
		Endpoint:         aws.String(global.Config.AwsS3.Endpoint), //Minio sets the address here, which is compatible with
		S3ForcePathStyle: aws.Bool(global.Config.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(global.Config.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			global.Config.AwsS3.SecretID,
			global.Config.AwsS3.SecretKey,
			"",
		),
	})
	return sess
}
