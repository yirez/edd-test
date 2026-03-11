package usecase

import (
	"net/http"

	"edd-test/cmd/notification-service/internal/repo"
)

type NotificationUseCase interface {
	HandleEvent(w http.ResponseWriter, r *http.Request)
}

type NotificationUseCaseImpl struct {
	deliveryRepo repo.DeliveryRepo
}

func NewNotificationUseCase(deliveryRepo repo.DeliveryRepo) NotificationUseCase {
	return &NotificationUseCaseImpl{
		deliveryRepo: deliveryRepo,
	}
}

// HandleEvent comment here
func (u *NotificationUseCaseImpl) HandleEvent(w http.ResponseWriter, _ *http.Request) {
	// Pattern: Event Notification - trigger side effects from business events.
	// Pattern: Idempotent Consumer - prevent duplicate notifications on re-delivery.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte(`{"message":"notification handler not implemented yet"}`))
}
