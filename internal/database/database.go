package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"new-mall/internal/global"
	"strings"
)

// NewDatabase initializes a new database connection.
func NewDatabase() *gorm.DB {
	mConfig := global.CONFIG.MySql
	dsn := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
