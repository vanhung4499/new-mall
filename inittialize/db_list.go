package initialize

import (
	"gorm.io/gorm"
	"new-mall/config"
	"new-mall/global"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.Config.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		case "pgsql":
			dbMap[info.AliasName] = GormPgSqlByConfig(config.Pgsql{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}
	// Make a special judgment to determine whether there is migration
	// Adapt to lower versions and migrate multiple database versions
	if sysDB, ok := dbMap[sys]; ok {
		global.DB = sysDB
	}
	global.DBList = dbMap
}
