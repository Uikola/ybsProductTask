package server

import (
	"database/sql"
	"github.com/Uikola/ybsProductTask/internal/db/repository/postgres"
	"github.com/Uikola/ybsProductTask/internal/server/courier"
	"github.com/Uikola/ybsProductTask/internal/server/order"
	"github.com/Uikola/ybsProductTask/internal/usecase"
	"github.com/go-chi/chi/v5"
	"log/slog"
)

func Router(db *sql.DB, router chi.Router, log *slog.Logger) {
	courierRepository := postgres.NewCourierRepository(db)
	orderRepository := postgres.NewOrderRepository(db)
	courierUseCase := usecase.NewCourierUC(courierRepository, orderRepository)
	orderUseCase := usecase.NewOrderUC(orderRepository)

	router.Route("/couriers", func(r chi.Router) {
		r.Post("/", courier.LoadCourier(courierUseCase, log))
		r.Get("/{courier_id}", courier.GetCourier(courierUseCase, log))
		r.Get("/", courier.GetCouriers(courierUseCase, log))
		r.Get("/meta-info/{courier_id}", courier.GetMetaInfo(courierUseCase, log))
	})
	router.Route("/orders", func(r chi.Router) {
		r.Post("/", order.LoadOrders(orderUseCase, log))
		r.Get("/{order_id}", order.GetOrder(orderUseCase, log))
		r.Get("/", order.GetOrders(orderUseCase, log))
		r.Put("/complete", order.CompleteOrder(orderUseCase, log))
	})
}
