package order

import (
	"encoding/json"
	"errors"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/errorz"
	sl "github.com/Uikola/ybsProductTask/internal/src/logger"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func CompleteOrder(useCase UseCase, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			CompleteInfo entity.CompleteOrderInfo `json:"complete_info"`
		}
		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Info("failed to parse the request", sl.Err(err))
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}

		order, err := useCase.CompleteOrder(ctx, request.CompleteInfo)
		if err != nil {
			if errors.Is(err, errorz.ErrOrderAlreadyExists) {
				log.Info("order already complete")
				http.Error(w, "order already complete", http.StatusBadRequest)
				return
			}
			log.Info("failed to complete order", sl.Err(err))
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, order)
	}
}
