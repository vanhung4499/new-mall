package upload

import (
	"errors"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"new-mall/global"
	"new-mall/utils"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct{}

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	//Read file suffix
	ext := path.Ext(file.Filename)
	// Read the file name and encrypt it
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// Splice new file name
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// Try to create this path
	mkdirErr := os.MkdirAll(global.Config.Local.Path, os.ModePerm)
	if mkdirErr != nil {
		global.Log.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// Concatenate path and file name
	p := global.Config.Local.Path + "/" + filename

	f, openError := file.Open() // read file
	if openError != nil {
		global.Log.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // Create file defer close

	out, createErr := os.Create(p)
	if createErr != nil {
		global.Log.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // Create file defer close

	_, copyErr := io.Copy(out, f) // Transfer (copy) files
	if copyErr != nil {
		global.Log.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description: Delete Files
//@param: key string
//@return: error

func (*Local) DeleteFile(key string) error {
	p := global.Config.Local.Path + "/" + key
	if strings.Contains(p, global.Config.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("Delete local file failed, err:" + err.Error())
		}
	}
	return nil
}
