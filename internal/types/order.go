package types

type CreateOrderReq struct {
	AddressID uint `form:"address_id" binding:"required"`
}
