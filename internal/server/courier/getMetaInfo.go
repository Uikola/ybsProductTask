package courier

import (
	sl "github.com/Uikola/ybsProductTask/internal/src/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func GetMetaInfo(useCase UseCase, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		courierID, err := strconv.Atoi(chi.URLParam(r, "courier_id"))
		if err != nil {
			log.Info("invalid courier id", sl.Err(err))
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		startDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("start_date"))
		if err != nil {
			log.Info("invalid start date", sl.Err(err))
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		endDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("end_date"))
		if err != nil {
			log.Info("invalid end date", sl.Err(err))
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		metaInfo, err := useCase.GetMetaInfo(ctx, courierID, startDate, endDate)
		if err != nil {
			log.Info("failed to get courier meta info", sl.Err(err))
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, metaInfo)
	}
}
