package order_usecase

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/errorz"
)

func (uc UseCaseImpl) CreateOrders(ctx context.Context, orders []entity.Order) error {
	return uc.repo.CreateOrders(ctx, orders)
}

func (uc UseCaseImpl) GetOrder(ctx context.Context, orderID int) (entity.Order, error) {
	return uc.repo.GetOrder(ctx, orderID)
}

func (uc UseCaseImpl) GetOrders(ctx context.Context, dto GetOrdersDTO) ([]entity.Order, error) {
	return uc.repo.GetOrders(ctx, dto.Offset, dto.Limit)
}

func (uc UseCaseImpl) CompleteOrder(ctx context.Context, completeInfo entity.CompleteOrderInfo) (int, error) {
	exists, err := uc.repo.Exists(ctx, completeInfo.OrderID)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errorz.ErrOrderAlreadyExists
	}

	return uc.repo.CompleteOrder(ctx, completeInfo)
}
