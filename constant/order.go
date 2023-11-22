package constant

const (
	OrderTypeUnPaid = iota + 1
	OrderTypePendingShipping
	OrderTypeShipping
	OrderTypeReceipt
)

var OrderTypeMap = map[int]string{
	OrderTypeUnPaid:          "Unpaid",
	OrderTypePendingShipping: "Paid, waiting for shipment",
	OrderTypeShipping:        "Shipped, waiting for receipt",
	OrderTypeReceipt:         "Goods received, transaction successful",
}
