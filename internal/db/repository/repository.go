package repository

import (
	"context"
	"errors"
	"github.com/Uikola/ybsProductTask/internal/entities"
)

var ErrCourierNotFound = errors.New("courier not found")
var ErrNoCouriers = errors.New("no couriers")

type CourierRepository interface {
	CreateCouriers(ctx context.Context, couriers []entities.Courier) error
	GetCourier(ctx context.Context, courierID int) (entities.Courier, error)
	GetCouriers(ctx context.Context, offset, limit int) ([]entities.Courier, error)
}

var ErrOrderNotFound = errors.New("order not found")
var ErrNoOrders = errors.New("no orders")

type OrderRepository interface {
	CreateOrders(ctx context.Context, orders []entities.Order) error
	GetOrder(ctx context.Context, orderID int) (entities.Order, error)
	GetOrders(ctx context.Context, offset, limit int) ([]entities.Order, error)
	CompleteOrder(ctx context.Context, completeInfo entities.CompleteOrderInfo) (int, error)
	GetOrdersByCourier(ctx context.Context, courierID int) ([]entities.Order, error)
	Exists(ctx context.Context, orderID int) (bool, error)
}
