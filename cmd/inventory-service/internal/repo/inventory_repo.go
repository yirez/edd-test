package repo

type InventoryRepo interface {
	Reserve(orderID string) error
	Release(orderID string) error
}

type InventoryRepoImpl struct{}

func NewInventoryRepo() InventoryRepo {
	return &InventoryRepoImpl{}
}

// Reserve comment here
func (r *InventoryRepoImpl) Reserve(orderID string) error {
	// todo add implementation
	return nil
}

// Release comment here
func (r *InventoryRepoImpl) Release(orderID string) error {
	// todo add implementation
	return nil
}
