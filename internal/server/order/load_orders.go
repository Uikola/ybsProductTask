package order

import (
	"encoding/json"
	"net/http"

	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/entity/types"
	"github.com/Uikola/ybsProductTask/internal/errorz"
)

func (h Handler) LoadOrders(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Orders []struct {
			Weight       int            `json:"weight"`
			Region       int            `json:"region"`
			DeliveryTime types.Interval `json:"delivery_time"`
			Price        int            `json:"price"`
		} `json:"orders"`
	}
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Error().Err(err).Msg("failed to parse the request")
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	orders := make([]entity.Order, len(request.Orders))
	for i, data := range request.Orders {
		order := entity.Order{
			Weight:       data.Weight,
			Region:       data.Region,
			DeliveryTime: []types.Interval{data.DeliveryTime},
			Price:        data.Price,
		}
		orders[i] = order
	}

	err = ValidateOrders(orders)
	if err != nil {
		h.log.Error().Err(err).Msg("failed to validate orders")
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.useCase.CreateOrders(ctx, orders)
	if err != nil {
		h.log.Error().Err(err).Msg("failed to save orders")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ValidateOrders(orders []entity.Order) error {
	for _, order := range orders {
		if order.Weight < 0 {
			return errorz.ErrInvalidWeight
		}
		if order.Price < 0 {
			return errorz.ErrInvalidPrice
		}
		if order.Region < 0 {
			return errorz.ErrInvalidRegion
		}
		if len(order.DeliveryTime) == 0 {
			return errorz.ErrInvalidDeliveryTime
		}
	}
	return nil
}
