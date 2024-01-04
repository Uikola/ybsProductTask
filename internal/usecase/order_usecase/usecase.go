package order_usecase

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/entity"
)

type Repo interface {
	CreateOrders(ctx context.Context, orders []entity.Order) error
	GetOrder(ctx context.Context, orderID int) (entity.Order, error)
	GetOrders(ctx context.Context, offset, limit int) ([]entity.Order, error)
	CompleteOrder(ctx context.Context, completeInfo entity.CompleteOrderInfo) (int, error)
	GetOrdersByCourier(ctx context.Context, courierID int) ([]entity.Order, error)
	Exists(ctx context.Context, orderID int) (bool, error)
}

type UseCaseImpl struct {
	repo Repo
}

func New(repo Repo) *UseCaseImpl {
	return &UseCaseImpl{repo: repo}
}
