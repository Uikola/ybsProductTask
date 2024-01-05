package courier

import (
	"encoding/json"
	"net/http"

	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/entity/types"
	"github.com/Uikola/ybsProductTask/internal/errorz"
	"github.com/rs/zerolog/log"
)

func (h Handler) LoadCouriers(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Couriers []struct {
			Type         entity.CourierType `json:"type"`
			Regions      []int              `json:"regions"`
			WorkingHours []types.Interval   `json:"working_hours"`
		} `json:"couriers"`
	}
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Error().Err(err).Msg("failed to parse the request")
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	couriers := make([]entity.Courier, len(request.Couriers))
	for i, data := range request.Couriers {
		courier := entity.Courier{
			Type:         data.Type,
			Regions:      data.Regions,
			WorkingHours: data.WorkingHours,
		}
		couriers[i] = courier
	}

	err = ValidateCouriers(couriers)
	if err != nil {
		h.log.Error().Err(err).Msg("failed to validate couriers")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.useCase.CreateCouriers(ctx, couriers)
	if err != nil {
		log.Error().Err(err).Msg("failed to save couriers")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
