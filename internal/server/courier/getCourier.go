package courier

import (
	"errors"

	"github.com/Uikola/ybsProductTask/internal/errorz"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"net/http"
	"strconv"
)

func (h Handler) GetCourier(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	courierID, err := strconv.Atoi(chi.URLParam(r, "courier_id"))
	if err != nil {
		h.log.Error().Err(err).Msg("invalid courier id")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	courier, err := h.useCase.GetCourier(ctx, courierID)
	switch {
	case errors.Is(err, errorz.ErrCourierNotFound):
		h.log.Error().Err(err).Msg("courier not found")
		http.Error(w, "courier not found", http.StatusNotFound)
		return
	case err != nil:
		h.log.Error().Err(err).Msg("failed to get a courier")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, courier)
}
