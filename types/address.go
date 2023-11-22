package types

type AddressServiceReq struct {
	Id      uint   `form:"id" json:"id"`
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

type AddressGetReq struct {
	Id uint `form:"id" json:"id"`
}

type AddressListReq struct {
	BasePage
}

type AddressDeleteReq struct {
	Id uint `form:"id" json:"id"`
}

type AddressCreateReq struct {
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

type AddressRes struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
}
