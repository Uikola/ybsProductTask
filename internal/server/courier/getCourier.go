package courier

import (
	"errors"
	"github.com/Uikola/ybsProductTask/internal/db/repository"
	sl "github.com/Uikola/ybsProductTask/internal/lib/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func GetCourier(repo repository.CourierRepository, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		courierID, err := strconv.Atoi(chi.URLParam(r, "courier_id"))
		if err != nil {
			log.Info("invalid courier id", sl.Err(err))
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		courier, err := repo.GetCourier(ctx, courierID)
		if err != nil {
			if errors.Is(err, repository.ErrCourierNotFound) {
				log.Info("courier not found", sl.Err(err))
				http.Error(w, "courier not found", http.StatusNotFound)
				return
			}
			log.Info("failed to get a courier", sl.Err(err))
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, courier)
	}
}
