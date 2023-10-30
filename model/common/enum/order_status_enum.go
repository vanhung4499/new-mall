package enum

type OrderStatusEnum int

const (
	DEFAULT                  OrderStatusEnum = -9
	ORDER_PRE_PAY            OrderStatusEnum = 0
	ORDER_PAID               OrderStatusEnum = 1
	ORDER_PACKAGED           OrderStatusEnum = 2
	ORDER_EXPRESS            OrderStatusEnum = 3
	ORDER_SUCCESS            OrderStatusEnum = 4
	ORDER_CLOSED_BY_MALLUSER OrderStatusEnum = -1
	ORDER_CLOSED_BY_EXPIRED  OrderStatusEnum = -2
	ORDER_CLOSED_BY_JUDGE    OrderStatusEnum = -3
)

func GetMallOrderStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "Pre Paid"
	case 1:
		return 1, "Paid"
	case 2:
		return 2, "Packaged"
	case 3:
		return 3, "Express"
	case 4:
		return 4, "Success"
	case -1:
		return -1, "Closed By mall user"
	case -2:
		return -2, "Closed By Expired"
	case -3:
		return -3, "Closed By Judge"
	default:
		return -9, "error"
	}
}

func (g OrderStatusEnum) Code() int {
	switch g {
	case ORDER_PRE_PAY:
		return 0
	case ORDER_PAID:
		return 1
	case ORDER_PACKAGED:
		return 2
	case ORDER_EXPRESS:
		return 3
	case ORDER_SUCCESS:
		return 4
	case ORDER_CLOSED_BY_MALLUSER:
		return -1
	case ORDER_CLOSED_BY_EXPIRED:
		return -2
	case ORDER_CLOSED_BY_JUDGE:
		return 3
	default:
		return -9
	}
}
