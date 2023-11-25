package global

import (
	"gorm.io/gorm"
	"new-mall/internal/config"
)

var (
	DB     *gorm.DB
	CONFIG *config.Server
)
