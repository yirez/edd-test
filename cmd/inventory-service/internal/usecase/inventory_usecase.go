package usecase

import (
	"net/http"

	"edd-test/cmd/inventory-service/internal/repo"
)

type InventoryUseCase interface {
	HandleOrderCreated(w http.ResponseWriter, r *http.Request)
}

type InventoryUseCaseImpl struct {
	inventoryRepo repo.InventoryRepo
}

func NewInventoryUseCase(inventoryRepo repo.InventoryRepo) InventoryUseCase {
	return &InventoryUseCaseImpl{
		inventoryRepo: inventoryRepo,
	}
}

// HandleOrderCreated comment here
func (u *InventoryUseCaseImpl) HandleOrderCreated(w http.ResponseWriter, _ *http.Request) {
	// Pattern: Idempotent Consumer - skip duplicate order.created deliveries before reserving stock.
	// Pattern: Event Notification - react to order.created without synchronous order-service calls.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte(`{"message":"inventory handler not implemented yet"}`))
}
