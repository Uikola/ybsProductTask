package entities

import (
	"github.com/Uikola/ybsProductTask/internal/entities/types"
)

type Courier struct {
	ID           int              `json:"courier_id"`
	Type         CourierType      `json:"courier_type"`
	Regions      []int            `json:"regions"`
	WorkingHours []types.Interval `json:"working_hours"`
}

type CourierType string

const (
	FootCourier CourierType = "FOOT"
	BikeCourier CourierType = "BIKE"
	AutoCourier CourierType = "AUTO"
)

type CourierMeta struct {
	Income int `json:"income"`
	Rating int `json:"rating"`
}
