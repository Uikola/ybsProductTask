package courier

import (
	"context"

	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/usecase/courier_usecase"
	"github.com/rs/zerolog"
)

type UseCase interface {
	CreateCouriers(ctx context.Context, couriers []entity.Courier) error
	GetCourier(ctx context.Context, courierID int) (entity.Courier, error)
	GetCouriers(ctx context.Context, dto courier_usecase.GetCouriersDTO) ([]entity.Courier, error)
	GetMetaInfo(ctx context.Context, dto courier_usecase.GetMetaInfoDTO) (entity.CourierMeta, error)
}

type Handler struct {
	useCase UseCase
	log     zerolog.Logger
}

func New(courierUseCase UseCase, log zerolog.Logger) Handler {
	return Handler{useCase: courierUseCase, log: log}
}
