package order

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Uikola/ybsProductTask/internal/errorz"
	"github.com/Uikola/ybsProductTask/internal/usecase/order_usecase"
	"github.com/go-chi/render"
)

func (h Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var offset, limit int
	ctx := r.Context()

	offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 1
	}
	dto := order_usecase.GetOrdersDTO{
		Offset: offset,
		Limit:  limit,
	}

	couriers, err := h.useCase.GetOrders(ctx, dto)
	switch {
	case errors.Is(err, errorz.ErrNoOrders):
		h.log.Error().Err(err).Msg("no orders")
		http.Error(w, "no orders", http.StatusNotFound)
		return
	case err != nil:
		h.log.Error().Err(err).Msg("failed to get orders")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, couriers)
}
