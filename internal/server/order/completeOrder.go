package order

import (
	"encoding/json"
	"github.com/Uikola/ybsProductTask/internal/db/repository"
	"github.com/Uikola/ybsProductTask/internal/entities"
	sl "github.com/Uikola/ybsProductTask/internal/lib/logger"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func CompleteOrder(repository repository.OrderRepository, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			CompleteInfo entities.CompleteOrderInfo `json:"complete_info"`
		}
		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Info("failed to parse the request", sl.Err(err))
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}

		exists, err := repository.Exists(ctx, request.CompleteInfo.OrderID)
		if err != nil {
			log.Info("bad request", sl.Err(err))
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		if exists {
			log.Info("order already complete")
			http.Error(w, "order already complete", http.StatusBadRequest)
			return
		}

		order, err := repository.CompleteOrder(ctx, request.CompleteInfo)
		if err != nil {
			log.Info("failed to complete order", sl.Err(err))
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, order)
	}
}
