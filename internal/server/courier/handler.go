package courier

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"time"
)

type UseCase interface {
	CreateCouriers(ctx context.Context, couriers []entity.Courier) error
	GetCourier(ctx context.Context, courierID int) (entity.Courier, error)
	GetCouriers(ctx context.Context, offset, limit int) ([]entity.Courier, error)
	GetMetaInfo(ctx context.Context, courierID int, startDate, endDate time.Time) (entity.CourierMeta, error)
}
