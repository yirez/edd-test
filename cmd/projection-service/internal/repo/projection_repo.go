package repo

type OrderProjectionRepo interface {
	UpsertOrderStatus(orderID, status string) error
}

type OrderProjectionRepoImpl struct{}

func NewOrderProjectionRepo() OrderProjectionRepo {
	return &OrderProjectionRepoImpl{}
}

// UpsertOrderStatus comment here
func (r *OrderProjectionRepoImpl) UpsertOrderStatus(orderID, status string) error {
	//todo add implementation
	return nil
}
