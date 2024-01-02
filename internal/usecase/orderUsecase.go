package usecase

import (
	"context"
	"errors"
	"github.com/Uikola/ybsProductTask/internal/db/repository"
	"github.com/Uikola/ybsProductTask/internal/entities"
)

var ErrOrderAlreadyExists = errors.New("order already exists")

type OrderUseCase interface {
	CreateOrders(ctx context.Context, orders []entities.Order) error
	GetOrder(ctx context.Context, orderID int) (entities.Order, error)
	GetOrders(ctx context.Context, offset, limit int) ([]entities.Order, error)
	CompleteOrder(ctx context.Context, completeInfo entities.CompleteOrderInfo) (int, error)
}

type OrderUC struct {
	repo repository.OrderRepository
}

func NewOrderUC(repo repository.OrderRepository) *OrderUC {
	return &OrderUC{repo: repo}
}

func (uc *OrderUC) CreateOrders(ctx context.Context, orders []entities.Order) error {
	return uc.repo.CreateOrders(ctx, orders)
}

func (uc *OrderUC) GetOrder(ctx context.Context, orderID int) (entities.Order, error) {
	return uc.repo.GetOrder(ctx, orderID)
}

func (uc *OrderUC) GetOrders(ctx context.Context, offset, limit int) ([]entities.Order, error) {
	return uc.repo.GetOrders(ctx, offset, limit)
}

func (uc *OrderUC) CompleteOrder(ctx context.Context, completeInfo entities.CompleteOrderInfo) (int, error) {
	exists, err := uc.repo.Exists(ctx, completeInfo.OrderID)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, ErrOrderAlreadyExists
	}

	return uc.repo.CompleteOrder(ctx, completeInfo)
}
