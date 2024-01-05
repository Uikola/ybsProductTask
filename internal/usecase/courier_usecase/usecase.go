package courier_usecase

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/entity/types"
	"time"
)

type LoadCouriersDTO struct {
	Type         entity.CourierType `json:"type"`
	Regions      []int              `json:"regions"`
	WorkingHours []types.Interval   `json:"working_hours"`
}

type GetMetaInfoDTO struct {
	CourierID int
	StartDate time.Time
	EndDate   time.Time
}

type GetCouriersDTO struct {
	Offset int
	Limit  int
}

type CourierRepo interface {
	CreateCouriers(ctx context.Context, couriers []entity.Courier) error
	GetCourier(ctx context.Context, courierID int) (entity.Courier, error)
	GetCouriers(ctx context.Context, offset, limit int) ([]entity.Courier, error)
}

type OrderRepo interface {
	CreateOrders(ctx context.Context, orders []entity.Order) error
	GetOrder(ctx context.Context, orderID int) (entity.Order, error)
	GetOrders(ctx context.Context, offset, limit int) ([]entity.Order, error)
	CompleteOrder(ctx context.Context, completeInfo entity.CompleteOrderInfo) (int, error)
	GetOrdersByCourier(ctx context.Context, courierID int) ([]entity.Order, error)
	Exists(ctx context.Context, orderID int) (bool, error)
}

type UseCaseImp struct {
	courierRepo CourierRepo
	orderRepo   OrderRepo
}

func New(courierRepo CourierRepo, orderRepo OrderRepo) *UseCaseImp {
	return &UseCaseImp{courierRepo: courierRepo, orderRepo: orderRepo}
}
