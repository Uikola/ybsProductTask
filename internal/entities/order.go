package entities

import (
	"github.com/Uikola/ybsProductTask/internal/pkg/types"
	"time"
)

type Order struct {
	ID           int              `json:"order_id"`
	Weight       int              `json:"weight"`
	Region       int              `json:"region"`
	DeliveryTime []types.Interval `json:"delivery_time"`
	Price        int              `json:"price"`
	CompleteTime *time.Time       `json:"complete_time,omitempty"`
	CourierID    *int             `json:"-"`
}

type CompleteOrderInfo struct {
	OrderID      int       `json:"order_id"`
	CourierID    int       `json:"courier_id"`
	CompleteTime time.Time `json:"complete_time"`
}
