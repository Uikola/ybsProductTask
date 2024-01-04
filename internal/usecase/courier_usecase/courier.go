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

func (uc UseCaseImp) GetCouriers(ctx context.Context, offset, limit int) ([]entity.Courier, error) {
	return uc.courierRepo.GetCouriers(ctx, offset, limit)
}

func (uc UseCaseImp) GetMetaInfo(ctx context.Context, courierID int, startDate, endDate time.Time) (entity.CourierMeta, error) {
	courier, err := uc.courierRepo.GetCourier(ctx, courierID)
	if err != nil {
		return entity.CourierMeta{}, err
	}

	orders, err := uc.orderRepo.GetOrdersByCourier(ctx, courierID)
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
		if InTimeSpan(startDate, endDate, *order.CompleteTime) {
			income += order.Price * incomeC
			ordersCount++
		}
	}

	hoursBetweenStartEnd := CalculateHours(endDate) - CalculateHours(startDate)
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
