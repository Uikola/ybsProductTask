package order

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Uikola/ybsProductTask/internal/errorz"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orderID, err := strconv.Atoi(chi.URLParam(r, "order_id"))
	if err != nil {
		h.log.Error().Err(err).Msg("invalid order id")
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	order, err := h.useCase.GetOrder(ctx, orderID)
	switch {
	case errors.Is(err, errorz.ErrOrderNotFound):
		h.log.Error().Err(err).Msg("order not found")
		http.Error(w, "order not found", http.StatusNotFound)
		return
	case err != nil:
		h.log.Error().Err(err).Msg("failed to get an order")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, order)
}
