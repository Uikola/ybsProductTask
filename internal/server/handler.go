package server

import (
	"database/sql"

	"github.com/Uikola/ybsProductTask/internal/db/repository/postgres"
	"github.com/Uikola/ybsProductTask/internal/server/courier"
	"github.com/Uikola/ybsProductTask/internal/server/order"
	"github.com/Uikola/ybsProductTask/internal/usecase/courier_usecase"
	"github.com/Uikola/ybsProductTask/internal/usecase/order_usecase"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type Handler struct {
	Courier courier.Handler
	Order   order.Handler
}

func New(courierHandler courier.Handler, orderHandler order.Handler) *Handler {
	return &Handler{Courier: courierHandler, Order: orderHandler}
}

func Router(db *sql.DB, router chi.Router, log zerolog.Logger) {
	courierRepository := postgres.NewCourierRepository(db)
	orderRepository := postgres.NewOrderRepository(db)
	courierUseCase := courier_usecase.New(courierRepository, orderRepository)
	orderUseCase := order_usecase.New(orderRepository)
	handler := New(courier.New(courierUseCase, log), order.New(orderUseCase, log))

	router.Route("/couriers", func(r chi.Router) {
		r.Post("/", handler.Courier.LoadCouriers)
		r.Get("/{courier_id}", handler.Courier.GetCourier)
		r.Get("/", handler.Courier.GetCouriers)
		r.Get("/meta-info/{courier_id}", handler.Courier.GetMetaInfo)
	})
	router.Route("/orders", func(r chi.Router) {
		r.Post("/", handler.Order.LoadOrders)
		r.Get("/{order_id}", handler.Order.GetOrder)
		r.Get("/", handler.Order.GetOrders)
		r.Put("/complete", handler.Order.CompleteOrder)
	})
}
