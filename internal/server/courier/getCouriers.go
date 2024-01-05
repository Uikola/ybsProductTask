package courier

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Uikola/ybsProductTask/internal/errorz"
	"github.com/Uikola/ybsProductTask/internal/usecase/courier_usecase"
	"github.com/go-chi/render"
)

func (h Handler) GetCouriers(w http.ResponseWriter, r *http.Request) {
	var offset, limit int
	ctx := r.Context()

	offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 1
	}
	dto := courier_usecase.GetCouriersDTO{
		Offset: offset,
		Limit:  limit,
	}

	couriers, err := h.useCase.GetCouriers(ctx, dto)
	switch {
	case errors.Is(err, errorz.ErrNoCouriers):
		h.log.Error().Err(err).Msg("no couriers")
		http.Error(w, "no couriers", http.StatusNotFound)
		return
	case err != nil:
		h.log.Error().Err(err).Msg("failed to get couriers")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, couriers)
}
