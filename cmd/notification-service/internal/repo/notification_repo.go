package repo

type DeliveryRepo interface {
	Record(messageID string) error
}

type DeliveryRepoImpl struct{}

func NewDeliveryRepo() DeliveryRepo {
	return &DeliveryRepoImpl{}
}

// Record comment here
func (r *DeliveryRepoImpl) Record(messageID string) error {
	// todo add implementation
	return nil
}
