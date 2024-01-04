package courier_usecase_test

import (
	"context"
	"errors"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/errorz"
	"github.com/Uikola/ybsProductTask/internal/usecase/courier_usecase"
	"github.com/Uikola/ybsProductTask/internal/usecase/courier_usecase/mocks"
	mocks2 "github.com/Uikola/ybsProductTask/internal/usecase/order_usecase/mocks"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetC(t *testing.T) {
	cases := []struct {
		name        string
		courierType entity.CourierType
		expIncomeC  int
		expRatingC  int
	}{
		{
			name:        "foot courier",
			courierType: entity.FootCourier,
			expIncomeC:  2,
			expRatingC:  3,
		},
		{
			name:        "bike courier",
			courierType: entity.BikeCourier,
			expIncomeC:  3,
			expRatingC:  2,
		},
		{
			name:        "auto courier",
			courierType: entity.AutoCourier,
			expIncomeC:  4,
			expRatingC:  1,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			t.Parallel()

			incomeC, ratingC := courier_usecase.GetC(tCase.courierType)
			require.Equal(t, tCase.expIncomeC, incomeC)
			require.Equal(t, tCase.expRatingC, ratingC)
		})
	}
}

func TestInTimeSpan(t *testing.T) {
	cases := []struct {
		name      string
		startDate time.Time
		endDate   time.Time
		date      time.Time
		expResp   bool
	}{
		{
			name:      "date is in time span",
			startDate: time.Date(2022, 10, 25, 12, 26, 27, 10, time.UTC),
			endDate:   time.Date(2023, 5, 25, 12, 26, 27, 10, time.UTC),
			date:      time.Date(2022, 12, 25, 12, 26, 27, 10, time.UTC),
			expResp:   true,
		},
		{
			name:      "date isn't in time span",
			startDate: time.Date(2022, 10, 25, 12, 26, 27, 10, time.UTC),
			endDate:   time.Date(2023, 5, 25, 12, 26, 27, 10, time.UTC),
			date:      time.Date(2019, 12, 25, 12, 26, 27, 10, time.UTC),
			expResp:   false,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			t.Parallel()

			ok := courier_usecase.InTimeSpan(tCase.startDate, tCase.endDate, tCase.date)
			require.Equal(t, tCase.expResp, ok)
		})
	}
}

func TestCalculateHours(t *testing.T) {
	date := time.Date(2007, 6, 5, 0, 20, 0, 0, time.UTC)
	hours := courier_usecase.CalculateHours(date)
	require.Equal(t, 17585760, hours)
}

func TestUseCaseImpl_GetMetaInfo(t *testing.T) {
	ctx := context.Background()
	repoErr := errors.New("repo err")
	courierID, startDate, endDate := 1, time.Date(2023, 12, 20, 0, 0, 0, 0, time.UTC), time.Date(2023, 12, 20, 1, 0, 0, 0, time.UTC)
	cases := []struct {
		name                   string
		getCourierResp         entity.Courier
		getOrdersByCourierResp []entity.Order
		getCourierErr          error
		getOrdersByCourierErr  error
		expResp                entity.CourierMeta
		expErr                 error
	}{
		{
			name:                   "success",
			getCourierResp:         entity.Courier{ID: 1, Type: "FOOT"},
			getOrdersByCourierResp: []entity.Order{{Price: 100, CompleteTime: getTime(time.Date(2023, 12, 20, 0, 20, 0, 0, time.UTC))}},
			expResp:                entity.CourierMeta{Income: 200, Rating: 3},
		},
		{
			name:          "get courier error",
			getCourierErr: errorz.ErrCourierNotFound,
			expErr:        errorz.ErrCourierNotFound,
		},
		{
			name:                  "get orders by courier error",
			getOrdersByCourierErr: repoErr,
			expResp:               entity.CourierMeta{},
			expErr:                repoErr,
		},
		{
			name:    "empty orders",
			expResp: entity.CourierMeta{},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			courierRepo := mocks.NewMockRepo(t)
			orderRepo := mocks2.NewMockRepo(t)
			courierRepo.EXPECT().GetCourier(ctx, courierID).Return(tCase.getCourierResp, tCase.getCourierErr)
			if tCase.getCourierErr == nil {
				orderRepo.EXPECT().GetOrdersByCourier(ctx, courierID).Return(tCase.getOrdersByCourierResp, tCase.getOrdersByCourierErr)
			}

			useCase := courier_usecase.New(courierRepo, orderRepo)
			metaInfo, err := useCase.GetMetaInfo(ctx, courierID, startDate, endDate)
			require.ErrorIs(t, err, tCase.expErr)
			require.Equal(t, tCase.expResp, metaInfo)
		})
	}

}

func getTime(date time.Time) *time.Time {
	return &date
}
