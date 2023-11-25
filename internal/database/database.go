package database

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"new-mall/internal/global"
	"strings"
)

// NewDatabase initializes a new database connection.
func NewDatabase() *gorm.DB {
	mConfig := global.CONFIG.MySql
	dsn := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")

	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: ormLogger,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
