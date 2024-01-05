package courier_usecase

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"time"
)

func (uc UseCaseImp) CreateCouriers(ctx context.Context, couriers []entity.Courier) error {
	return uc.courierRepo.CreateCouriers(ctx, couriers)
}

func (uc UseCaseImp) GetCourier(ctx context.Context, courierID int) (entity.Courier, error) {
	return uc.courierRepo.GetCourier(ctx, courierID)
}

func (uc UseCaseImp) GetCouriers(ctx context.Context, dto GetCouriersDTO) ([]entity.Courier, error) {
	return uc.courierRepo.GetCouriers(ctx, dto.Offset, dto.Limit)
}

func (uc UseCaseImp) GetMetaInfo(ctx context.Context, dto GetMetaInfoDTO) (entity.CourierMeta, error) {
	courier, err := uc.courierRepo.GetCourier(ctx, dto.CourierID)
	if err != nil {
		return entity.CourierMeta{}, err
	}

	orders, err := uc.orderRepo.GetOrdersByCourier(ctx, dto.CourierID)
	if err != nil {
		return entity.CourierMeta{}, err
	}
	if orders == nil {
		return entity.CourierMeta{}, nil
	}
	var income, rating, ordersCount int

	// c is the coefficient that we should use to calculate income and rating
	incomeC, ratingC := GetC(courier.Type)

	for _, order := range orders {
		if InTimeSpan(dto.StartDate, dto.EndDate, *order.CompleteTime) {
			income += order.Price * incomeC
			ordersCount++
		}
	}

	hoursBetweenStartEnd := CalculateHours(dto.EndDate) - CalculateHours(dto.StartDate)
	rating = (ordersCount / hoursBetweenStartEnd) * ratingC

	return entity.CourierMeta{Income: income, Rating: rating}, nil
}

func GetC(courierType entity.CourierType) (int, int) {
	var incomeC, ratingC int

	switch courierType {
	case entity.FootCourier:
		incomeC = 2
		ratingC = 3
	case entity.BikeCourier:
		incomeC = 3
		ratingC = 2
	case entity.AutoCourier:
		incomeC = 4
		ratingC = 1
	}
	return incomeC, ratingC
}

func InTimeSpan(start, end, date time.Time) bool {
	if start.Before(end) {
		return !date.Before(start) && !date.After(end)
	}
	if start.Equal(end) {
		return date.Equal(start)
	}
	return !start.After(date) || !end.Before(date)
}

func CalculateHours(date time.Time) int {
	return date.Year()*365*24 + int(date.Month())*30*24 + date.Day()*24 + date.Hour()
}
