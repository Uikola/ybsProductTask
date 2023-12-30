package order

import (
	"encoding/json"
	"errors"
	"github.com/Uikola/ybsProductTask/internal/db/repository"
	"github.com/Uikola/ybsProductTask/internal/entities"
	sl "github.com/Uikola/ybsProductTask/internal/lib/logger"
	"github.com/Uikola/ybsProductTask/internal/pkg/types"
	"log/slog"
	"net/http"
)

var ErrInvalidWeight = errors.New("invalid weight")
var ErrInvalidPrice = errors.New("invalid price")
var ErrInvalidRegion = errors.New("invalid region")

func LoadOrders(repository repository.OrderRepository, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			log.Info("failed to parse the request", sl.Err(err))
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}

		var orders []entities.Order

		for _, data := range request.Orders {
			order := entities.Order{
				Weight:       data.Weight,
				Region:       data.Region,
				DeliveryTime: []types.Interval{data.DeliveryTime},
				Price:        data.Price,
			}
			orders = append(orders, order)
		}

		err = validateOrders(orders)
		if err != nil {
			log.Info("failed to validate orders", sl.Err(err))
			http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = repository.CreateOrders(ctx, orders)
		if err != nil {
			log.Info("failed to save orders", sl.Err(err))
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func validateOrders(orders []entities.Order) error {
	for _, order := range orders {
		if order.Weight < 0 {
			return ErrInvalidWeight
		}
		if order.Price < 0 {
			return ErrInvalidPrice
		}
		if order.Region < 0 {
			return ErrInvalidRegion
		}
	}
	return nil
}
