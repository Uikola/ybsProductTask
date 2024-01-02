package order

import (
	"errors"
	"github.com/Uikola/ybsProductTask/internal/db/repository"
	sl "github.com/Uikola/ybsProductTask/internal/src/logger"
	"github.com/Uikola/ybsProductTask/internal/usecase"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func GetOrders(useCase usecase.OrderUseCase, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var offset, limit int
		ctx := r.Context()

		offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 1
		}

		couriers, err := useCase.GetOrders(ctx, offset, limit)
		if err != nil {
			if errors.Is(err, repository.ErrNoOrders) {
				log.Info("no orders", sl.Err(err))
				http.Error(w, "no orders", http.StatusNotFound)
				return
			}
			log.Info("failed to get orders", sl.Err(err))
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, couriers)
	}
}
