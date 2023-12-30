package server

import (
	"database/sql"
	"github.com/Uikola/ybsProductTask/internal/db/repository/postgres"
	"github.com/Uikola/ybsProductTask/internal/server/courier"
	"github.com/Uikola/ybsProductTask/internal/server/order"
	"github.com/go-chi/chi/v5"
	"log/slog"
)

func Router(db *sql.DB, router chi.Router, log *slog.Logger) {
	courierRepository := postgres.NewCourierRepository(db)
	orderRepository := postgres.NewOrderRepository(db)

	router.Route("/api", func(r chi.Router) {
		r.Post("/couriers", courier.LoadCourier(courierRepository, log))
		r.Get("/couriers/{courier_id}", courier.GetCourier(courierRepository, log))
		r.Get("/couriers", courier.GetCouriers(courierRepository, log))
		r.Post("/orders", order.LoadOrders(orderRepository, log))
		r.Get("/orders/{order_id}", order.GetOrder(orderRepository, log))
		r.Get("/orders", order.GetOrders(orderRepository, log))
		r.Put("/orders/complete", order.CompleteOrder(orderRepository, log))
	})
}
