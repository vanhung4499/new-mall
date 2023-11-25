package uploadprovider

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"new-mall/internal/global"
	"new-mall/pkg/component/hasher"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct{}

// UploadFile uploads a file to the local file system
func (*Local) UploadFile(ctx context.Context, file *multipart.FileHeader) (string, string, error) {
	// Read file extension
	ext := path.Ext(file.Filename)
	// Read file name and encrypt it (MD5)
	name := strings.TrimSuffix(file.Filename, ext)
	name = hasher.Hash(name)
	// Concatenate the new file name
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// Attempt to create the storage path
	mkdirErr := os.MkdirAll(global.CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		return "", "", errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
	}
	// Concatenate the path and file name
	storePath := global.CONFIG.Local.StorePath + "/" + filename
	filepath := global.CONFIG.Local.Path + "/" + filename

	f, openError := file.Open() // Read the file
	if openError != nil {
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // Create file defer close

	out, createErr := os.Create(storePath)
	if createErr != nil {
		return "", "", errors.New("function os.Create() failed, err:" + createErr.Error())
	}
	defer out.Close() // Create file defer close

	_, copyErr := io.Copy(out, f) // Transfer (copy) the file
	if copyErr != nil {
		return "", "", errors.New("function io.Copy() failed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

// DeleteFile deletes a file from the local file system
func (*Local) DeleteFile(ctx context.Context, key string) error {
	p := global.CONFIG.Local.StorePath + "/" + key
	if strings.Contains(p, global.CONFIG.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("Local file deletion failed, err:" + err.Error())
		}
	}
	return nil
}
