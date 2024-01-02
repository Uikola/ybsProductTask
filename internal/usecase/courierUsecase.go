package usecase

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/db/repository"
	"github.com/Uikola/ybsProductTask/internal/entities"
	"time"
)

type CourierUseCase interface {
	CreateCouriers(ctx context.Context, couriers []entities.Courier) error
	GetCourier(ctx context.Context, courierID int) (entities.Courier, error)
	GetCouriers(ctx context.Context, offset, limit int) ([]entities.Courier, error)
	GetMetaInfo(ctx context.Context, courierID int, startDate, endDate time.Time) (entities.CourierMeta, error)
}

type CourierUC struct {
	courierRepo repository.CourierRepository
	orderRepo   repository.OrderRepository
}

func NewCourierUC(courierRepo repository.CourierRepository, orderRepo repository.OrderRepository) *CourierUC {
	return &CourierUC{courierRepo: courierRepo, orderRepo: orderRepo}
}

func (uc *CourierUC) CreateCouriers(ctx context.Context, couriers []entities.Courier) error {
	return uc.courierRepo.CreateCouriers(ctx, couriers)
}

func (uc *CourierUC) GetCourier(ctx context.Context, courierID int) (entities.Courier, error) {
	return uc.courierRepo.GetCourier(ctx, courierID)
}

func (uc *CourierUC) GetCouriers(ctx context.Context, offset, limit int) ([]entities.Courier, error) {
	return uc.courierRepo.GetCouriers(ctx, offset, limit)
}

func (uc *CourierUC) GetMetaInfo(ctx context.Context, courierID int, startDate, endDate time.Time) (entities.CourierMeta, error) {
	courier, err := uc.courierRepo.GetCourier(ctx, courierID)
	if err != nil {
		return entities.CourierMeta{}, err
	}

	orders, err := uc.orderRepo.GetOrdersByCourier(ctx, courierID)
	if err != nil {
		return entities.CourierMeta{}, err
	}
	if orders == nil {
		return entities.CourierMeta{}, nil
	}
	var income, rating, ordersCount int

	incomeC, ratingC := getC(courier.Type)

	for _, order := range orders {
		if inTimeSpan(startDate, endDate, *order.CompleteTime) {
			income += order.Price * incomeC
			ordersCount++
		}
	}
	startHours := startDate.Year()*365*24 + int(startDate.Month())*30*24 + startDate.Day()*24 + startDate.Hour()
	endHours := endDate.Year()*365*24 + int(endDate.Month())*30*24 + endDate.Day()*24 + endDate.Hour()
	hoursBetweenStartEnd := endHours - startHours
	rating = (ordersCount / hoursBetweenStartEnd) * ratingC

	return entities.CourierMeta{Income: income, Rating: rating}, nil
}

func getC(courierType entities.CourierType) (int, int) {
	var incomeC, ratingC int

	switch courierType {
	case entities.FootCourier:
		incomeC = 2
		ratingC = 3
	case entities.BikeCourier:
		incomeC = 3
		ratingC = 2
	case entities.AutoCourier:
		incomeC = 4
		ratingC = 1
	}
	return incomeC, ratingC
}

func inTimeSpan(start, end, date time.Time) bool {
	if start.Before(end) {
		return !date.Before(start) && !date.After(end)
	}
	if start.Equal(end) {
		return date.Equal(start)
	}
	return !start.After(date) || !end.Before(date)
}
