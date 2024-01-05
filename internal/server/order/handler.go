package order

import (
	"context"

	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/usecase/order_usecase"
	"github.com/rs/zerolog"
)

type UseCase interface {
	CreateOrders(ctx context.Context, orders []entity.Order) error
	GetOrder(ctx context.Context, orderID int) (entity.Order, error)
	GetOrders(ctx context.Context, dto order_usecase.GetOrdersDTO) ([]entity.Order, error)
	CompleteOrder(ctx context.Context, completeInfo entity.CompleteOrderInfo) (int, error)
}

type Handler struct {
	useCase UseCase
	log     zerolog.Logger
}

func New(orderUseCase UseCase, log zerolog.Logger) Handler {
	return Handler{useCase: orderUseCase, log: log}
}
