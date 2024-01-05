package order

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/errorz"
	"github.com/go-chi/render"
)

func (h Handler) CompleteOrder(w http.ResponseWriter, r *http.Request) {
	var request struct {
		CompleteInfo entity.CompleteOrderInfo `json:"complete_info"`
	}
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.log.Error().Err(err).Msg("failed to parse the request")
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	order, err := h.useCase.CompleteOrder(ctx, request.CompleteInfo)
	switch {
	case errors.Is(err, errorz.ErrOrderAlreadyExists):
		h.log.Error().Err(err).Msg("order already complete")
		http.Error(w, "order already complete", http.StatusBadRequest)
		return
	case err != nil:
		h.log.Error().Err(err).Msg("failed to complete order")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, order)
}
