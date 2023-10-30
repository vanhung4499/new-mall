package internal

import (
	"fmt"
	"gorm.io/gorm/logger"
	"new-mall/global"
)

type writer struct {
	logger.Writer
}

// NewWriter writer Constructor
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf Format print log
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.Config.System.DbType {
	case "mysql":
		logZap = global.Config.Mysql.LogZap
	case "pgsql":
		logZap = global.Config.Pgsql.LogZap
	}
	if logZap {
		global.Log.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
