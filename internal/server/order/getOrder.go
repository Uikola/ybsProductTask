package order

import (
	"errors"
	"github.com/Uikola/ybsProductTask/internal/db/repository"
	sl "github.com/Uikola/ybsProductTask/internal/src/logger"
	"github.com/Uikola/ybsProductTask/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func GetOrder(useCase usecase.OrderUseCase, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		orderID, err := strconv.Atoi(chi.URLParam(r, "order_id"))
		if err != nil {
			log.Info("invalid order id", sl.Err(err))
			http.Error(w, "invalid order id", http.StatusBadRequest)
			return
		}

		order, err := useCase.GetOrder(ctx, orderID)
		if err != nil {
			if errors.Is(err, repository.ErrOrderNotFound) {
				log.Info("order not found", sl.Err(err))
				http.Error(w, "order not found", http.StatusNotFound)
				return
			}
			log.Info("failed to get an order", sl.Err(err))
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, order)
	}
}
