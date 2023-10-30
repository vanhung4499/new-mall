package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"new-mall/global"
	"os"
)

// Gorm initializes the database and generates database global variables
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables is dedicated to registration database tables
// Author SliverHorn
func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
	// System module table

	)
	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Log.Info("register table success")
}
