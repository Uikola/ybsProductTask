package entity

import (
	"github.com/Uikola/ybsProductTask/internal/entity/types"
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

func (e CourierType) Valid() bool {
	switch e {
	case FootCourier,
		BikeCourier,
		AutoCourier:
		return true
	}
	return false
}

type CourierMeta struct {
	Income int `json:"income"`
	Rating int `json:"rating"`
}
