package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"new-mall/config"
	"new-mall/pkg/utils"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct{}

// UploadFile uploads a file to the local file system
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// Read file extension
	ext := path.Ext(file.Filename)
	// Read file name and encrypt it (MD5)
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// Concatenate the new file name
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// Attempt to create the storage path
	mkdirErr := os.MkdirAll(config.Config.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		utils.Logger.Error("function os.MkdirAll() failed", mkdirErr.Error())
		return "", "", errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
	}
	// Concatenate the path and file name
	p := config.Config.Local.StorePath + "/" + filename
	filepath := config.Config.Local.Path + "/" + filename

	f, openError := file.Open() // Read the file
	if openError != nil {
		utils.Logger.Error("function file.Open() failed", openError.Error())
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // Create file defer close

	out, createErr := os.Create(p)
	if createErr != nil {
		utils.Logger.Error("function os.Create() failed", createErr.Error())

		return "", "", errors.New("function os.Create() failed, err:" + createErr.Error())
	}
	defer out.Close() // Create file defer close

	_, copyErr := io.Copy(out, f) // Transfer (copy) the file
	if copyErr != nil {
		utils.Logger.Error("function io.Copy() failed", copyErr.Error())
		return "", "", errors.New("function io.Copy() failed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

// DeleteFile deletes a file from the local file system
func (*Local) DeleteFile(key string) error {
	p := config.Config.Local.StorePath + "/" + key
	if strings.Contains(p, config.Config.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("Local file deletion failed, err:" + err.Error())
		}
	}
	return nil
}
