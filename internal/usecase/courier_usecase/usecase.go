package courier_usecase

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/usecase/order_usecase"
)

type Repo interface {
	CreateCouriers(ctx context.Context, couriers []entity.Courier) error
	GetCourier(ctx context.Context, courierID int) (entity.Courier, error)
	GetCouriers(ctx context.Context, offset, limit int) ([]entity.Courier, error)
}

type UseCaseImp struct {
	courierRepo Repo
	orderRepo   order_usecase.Repo // TODO разобраться куда положить репозиторий
}

func New(courierRepo Repo, orderRepo order_usecase.Repo) *UseCaseImp {
	return &UseCaseImp{courierRepo: courierRepo, orderRepo: orderRepo}
}
