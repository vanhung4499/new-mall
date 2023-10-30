package utils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"new-mall/global"
	"os"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetWriteSyncer
//@description: zap logger中加入file-rotatelogs
//@return: zapcore.WriteSyncer, error

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, //The location of the log file
		MaxSize:    10,   //The maximum size of the log file in MB before cutting is done
		MaxBackups: 200,  //The maximum number of old files to keep
		MaxAge:     30,   //Maximum number of days to keep old files
		Compress:   true, //Whether to compress/archive old files
	}

	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
