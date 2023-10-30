package enum

type ProductStatusEnum int

const (
	PRODUCT_DEFAULT ProductStatusEnum = -9
	PRODUCT_UNDER   ProductStatusEnum = 0
)

func GetProductStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "removed"
	default:
		return -9, "error"
	}
}

func (g ProductStatusEnum) Code() int {
	switch g {
	case PRODUCT_UNDER:
		return 0
	default:
		return -9
	}
}
