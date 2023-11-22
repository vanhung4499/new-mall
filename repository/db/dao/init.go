package dao

import (
	"context"
	"new-mall/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var (
	_db *gorm.DB
)

func InitMySQL() {
	mConfig := config.Config.MySql["default"]
	pathRead := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")
	pathWrite := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")

	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead, // DSN data source name
		DefaultStringSize:         256,      // default length for string-types fields
		DisableDatetimePrecision:  true,     // disable datetime precision, not supported in MySQL 5.6 and earlier
		DontSupportRenameIndex:    true,     // use delete and create for renaming indexes, not supported in MySQL 5.7 and MariaDB
		DontSupportRenameColumn:   true,     // use `change` for renaming columns, not supported in MySQL 8 and MariaDB
		SkipInitializeWithVersion: false,    // automatically configure based on version
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // set connection pool for idle connections
	sqlDB.SetMaxOpenConns(100) // set maximum open connections
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			// `db2` as sources, `db3` and `db4` as replicas
			Sources:  []gorm.Dialector{mysql.Open(pathRead)},                         // write operation
			Replicas: []gorm.Dialector{mysql.Open(pathWrite), mysql.Open(pathWrite)}, // read operation
			Policy:   dbresolver.RandomPolicy{},                                      // sources/replicas load balancing policy
		}))

	_db = _db.Set("gorm:table_options", "charset=utf8mb4")
	err = migrate()
	if err != nil {
		panic(err)
	}
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
