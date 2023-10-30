package global

import (
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"new-mall/config"
	"sync"
)

var (
	DB     *gorm.DB
	DBList map[string]*gorm.DB
	Redis  *redis.Client
	Mongo  *qmgo.QmgoClient
	Config config.Server
	Viper  *viper.Viper
	Log    *zap.Logger

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName Gets the db in the db list by name
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName Gets db by name and panics if it does not exist
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
