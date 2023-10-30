package model

import (
	"time"
)

// Token Structure
type Token struct {
	UserId     int       `json:"userId" form:"userId" gorm:"primarykey;AUTO_INCREMENT"`
	Token      string    `json:"token" form:"token" gorm:"column:token;type:varchar(32);"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;type:datetime"`
	ExpireTime time.Time `json:"expireTime" form:"expireTime" gorm:"column:expire_time;type:datetime"`
}

// TableName Token
func (Token) TableName() string {
	return "user_tokens"
}
