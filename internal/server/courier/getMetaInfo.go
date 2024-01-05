package courier

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Uikola/ybsProductTask/internal/usecase/courier_usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h Handler) GetMetaInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	courierID, err := strconv.Atoi(chi.URLParam(r, "courier_id"))
	if err != nil {
		h.log.Error().Err(err).Msg("invalid courier id")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("start_date"))
	if err != nil {
		h.log.Error().Err(err).Msg("invalid start date")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("end_date"))
	if err != nil {
		h.log.Error().Err(err).Msg("invalid end date")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	dto := courier_usecase.GetMetaInfoDTO{
		CourierID: courierID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	metaInfo, err := h.useCase.GetMetaInfo(ctx, dto)
	if err != nil {
		h.log.Error().Err(err).Msg("failed to get courier meta info")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, metaInfo)
}
