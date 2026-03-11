package events

const (
	OrderCreatedType      = "order.created"
	InventoryReservedType = "inventory.reserved"
	InventoryFailedType   = "inventory.failed"
	PaymentAuthorizedType = "payment.authorized"
	PaymentFailedType     = "payment.failed"
	OrderConfirmedType    = "order.confirmed"
	OrderRejectedType     = "order.rejected"
)

type OrderCreated struct {
	OrderID     string `json:"order_id"`
	CustomerID  string `json:"customer_id"`
	TotalAmount int64  `json:"total_amount"`
}
