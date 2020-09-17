package __struct

import "time"

type DeliveryService struct{}

func (d DeliveryService) IsDeliveryValid(delivery Delivery) bool {
	return delivery.Date.After(time.Now().Truncate(time.Hour).AddDate(0, 0, 2))
}

type Delivery struct {
	Date time.Time
}
