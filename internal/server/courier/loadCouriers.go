package courier

import (
	"encoding/json"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/entity/types"
	"github.com/Uikola/ybsProductTask/internal/errorz"
	sl "github.com/Uikola/ybsProductTask/internal/src/logger"
	"log/slog"
	"net/http"
)

func LoadCouriers(useCase UseCase, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Couriers []struct {
				Type         entity.CourierType `json:"type"`
				Regions      []int              `json:"regions"`
				WorkingHours []types.Interval   `json:"working_hours"`
			} `json:"couriers"`
		}
		var couriers []entity.Courier
		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Info("failed to parse the request", sl.Err(err))
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}

		for _, data := range request.Couriers {
			courier := entity.Courier{
				Type:         data.Type,
				Regions:      data.Regions,
				WorkingHours: data.WorkingHours,
			}
			couriers = append(couriers, courier)
		}

		err = ValidateCouriers(couriers)
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

func ValidateCouriers(couriers []entity.Courier) error {
	for _, courier := range couriers {
		err := ValidateCourier(courier)
		return err
	}
	return nil
}

func ValidateCourier(courier entity.Courier) error {
	if !courier.Type.Valid() {
		return errorz.ErrInvalidCourierType
	}
	if len(courier.Regions) == 0 {
		return errorz.ErrInvalidRegion
	}
	for _, region := range courier.Regions {
		if region < 0 {
			return errorz.ErrInvalidRegion
		}
	}
	if len(courier.WorkingHours) == 0 {
		return errorz.ErrInvalidWorkingHours
	}
	return nil
}
