package order

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/entity"
)

type UseCase interface {
	CreateOrders(ctx context.Context, orders []entity.Order) error
	GetOrder(ctx context.Context, orderID int) (entity.Order, error)
	GetOrders(ctx context.Context, offset, limit int) ([]entity.Order, error)
	CompleteOrder(ctx context.Context, completeInfo entity.CompleteOrderInfo) (int, error)
}
