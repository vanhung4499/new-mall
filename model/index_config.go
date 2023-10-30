package model

import "new-mall/global"

type IndexConfig struct {
	global.Model
	ConfigName  string `json:"configName" form:"configName" gorm:"column:config_name;type:varchar(50);"`
	ConfigType  int    `json:"configType" form:"configType" gorm:"column:config_type;type:tinyint"`
	ProductId   int    `json:"productId" form:"productId" gorm:"column:productId;type:bigint"`
	RedirectUrl string `json:"redirectUrl" form:"redirectUrl" gorm:"column:redirect_url;type:varchar(100);"`
	ConfigRank  int    `json:"configRank" form:"configRank" gorm:"column:config_rank;type:int"`
	CreateUser  int    `json:"createUser" form:"createUser" gorm:"column:create_user;type:int"`
	UpdateUser  int    `json:"updateUser" form:"updateUser" gorm:"column:update_user;type:int"`
}

// TableName IndexConfig
func (IndexConfig) TableName() string {
	return "index_configs"
}
