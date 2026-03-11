package order

type Status string

const (
	StatusPending   Status = "pending"
	StatusConfirmed Status = "confirmed"
	StatusRejected  Status = "rejected"
	StatusCancelled Status = "cancelled"
)

type Order struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Status     Status `json:"status"`
}
