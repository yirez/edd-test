package usecase

import (
	"net/http"

	"edd-test/cmd/projection-service/internal/repo"
)

type ProjectionUseCase interface {
	ApplyEvent(w http.ResponseWriter, r *http.Request)
	GetOrderView(w http.ResponseWriter, r *http.Request)
}

type ProjectionUseCaseImpl struct {
	orderProjectionRepo repo.OrderProjectionRepo
}

func NewProjectionUseCase(orderProjectionRepo repo.OrderProjectionRepo) ProjectionUseCase {
	return &ProjectionUseCaseImpl{
		orderProjectionRepo: orderProjectionRepo,
	}
}

// ApplyEvent comment here
func (u *ProjectionUseCaseImpl) ApplyEvent(w http.ResponseWriter, _ *http.Request) {
	// Pattern: CQRS Projection - build read models from the event stream instead of write tables.
	// Pattern: Idempotent Consumer - ignore duplicate events when updating projections.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte(`{"message":"projection worker not implemented yet"}`))
}

// GetOrderView comment here
func (u *ProjectionUseCaseImpl) GetOrderView(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte(`{"message":"order projection not implemented yet"}`))
}
