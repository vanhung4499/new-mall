package model

type AdminUser struct {
	AdminUserId   int    `json:"adminUserId" form:"adminUserId" gorm:"primarykey;AUTO_INCREMENT"`
	LoginUserName string `json:"loginUserName" form:"loginUserName" gorm:"column:login_user_name;type:varchar(50);"`
	LoginPassword string `json:"loginPassword" form:"loginPassword" gorm:"column:login_password;type:varchar(50);"`
	NickName      string `json:"nickName" form:"nickName" gorm:"column:nick_name;type:varchar(50);"`
	Locked        int    `json:"locked" form:"locked" gorm:"column:locked;type:tinyint"`
}

func (AdminUser) TableName() string {
	return "admin_users"
}
