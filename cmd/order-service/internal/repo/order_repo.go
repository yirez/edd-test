package repo

import (
	"edd-test/pkg/common_types/order"
)

type OrderRepo interface {
	Save(value order.Order) error
	FindByID(id string) (order.Order, error)
}

type OrderRepoImpl struct{}

func NewOrderRepo() OrderRepo {
	return &OrderRepoImpl{}
}

// Save comment here
func (r *OrderRepoImpl) Save(value order.Order) error {
	// todo add implementation
	return nil
}

// FindByID comment here
func (r *OrderRepoImpl) FindByID(id string) (order.Order, error) {
	// todo add implementation
	return order.Order{}, nil
}
