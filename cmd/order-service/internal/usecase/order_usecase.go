package usecase

import (
	"net/http"

	"edd-test/cmd/order-service/internal/repo"
)

type OrderUseCase interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
	CancelOrder(w http.ResponseWriter, r *http.Request)
}

type OrderUseCaseImpl struct {
	orderRepo repo.OrderRepo
}

func NewOrderUseCase(orderRepo repo.OrderRepo) OrderUseCase {
	return &OrderUseCaseImpl{
		orderRepo: orderRepo,
	}
}

// CreateOrder comment here
func (u *OrderUseCaseImpl) CreateOrder(w http.ResponseWriter, _ *http.Request) {
	// Pattern: Transactional Outbox - create order state and queue order.created in one DB tx.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte(`{"message":"create order not implemented yet"}`))
}

// GetOrder comment here
func (u *OrderUseCaseImpl) GetOrder(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte(`{"message":"get order projection not implemented yet"}`))
}

// CancelOrder comment here
func (u *OrderUseCaseImpl) CancelOrder(w http.ResponseWriter, _ *http.Request) {
	// Pattern: Transactional Outbox - persist cancellation and queue order.cancelled atomically.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte(`{"message":"cancel order not implemented yet"}`))
}
