package utils

import (
	"go.uber.org/zap"
	"new-mall/global"
	"os"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: PathExists
//@description: Check the file directory exist
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDir
//@description: Create folders in batches
//@param: dirs ...string
//@return: err error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.Log.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.Log.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
