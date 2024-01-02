package courier

import (
	"encoding/json"
	"errors"
	"github.com/Uikola/ybsProductTask/internal/entities"
	"github.com/Uikola/ybsProductTask/internal/entities/types"
	sl "github.com/Uikola/ybsProductTask/internal/src/logger"
	"github.com/Uikola/ybsProductTask/internal/usecase"
	"log/slog"
	"net/http"
)

var ErrInvalidCourierType = errors.New("invalid courier type")
var ErrInvalidRegion = errors.New("invalid region")

func LoadCourier(useCase usecase.CourierUseCase, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Couriers []struct {
				Type         entities.CourierType `json:"type"`
				Regions      []int                `json:"regions"`
				WorkingHours []types.Interval     `json:"working_hours"`
			} `json:"couriers"`
		}
		var couriers []entities.Courier
		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Info("failed to parse the request", sl.Err(err))
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}

		for _, data := range request.Couriers {
			courier := entities.Courier{
				Type:         data.Type,
				Regions:      data.Regions,
				WorkingHours: data.WorkingHours,
			}
			couriers = append(couriers, courier)
		}

		err = validateCouriers(couriers)
		if err != nil {
			log.Info("failed to validate couriers", sl.Err(err))
			http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = useCase.CreateCouriers(ctx, couriers)
		if err != nil {
			log.Info("failed to save couriers", sl.Err(err))
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func validateCouriers(couriers []entities.Courier) error {
	for _, courier := range couriers {
		if courier.Type != entities.FootCourier && courier.Type != entities.BikeCourier && courier.Type != entities.AutoCourier {
			return ErrInvalidCourierType
		}
		for _, region := range courier.Regions {
			if region < 0 {
				return ErrInvalidRegion
			}
		}
	}
	return nil
}
