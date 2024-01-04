package order_usecase_test

import (
	"context"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/errorz"
	"github.com/Uikola/ybsProductTask/internal/usecase/order_usecase"
	"github.com/Uikola/ybsProductTask/internal/usecase/order_usecase/mocks"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUseCaseImpl_CompleteOrder(t *testing.T) {
	completeInfo := entity.CompleteOrderInfo{
		OrderID:      1,
		CourierID:    1,
		CompleteTime: time.Now(),
	}
	ctx := context.Background()
	cases := []struct {
		name             string
		repoExistsResp   bool
		repoCompleteResp int
		expResponse      int
		expErr           error
		repoErr          error
	}{
		{
			name:             "success",
			repoExistsResp:   false,
			repoCompleteResp: 1,
			expResponse:      1,
		},
		{
			name:           "error",
			repoExistsResp: true,
			expErr:         errorz.ErrOrderAlreadyExists,
		},
		{
			name:           "repo_error",
			repoExistsResp: false,
			expErr:         errorz.ErrOrderNotFound,
			repoErr:        errorz.ErrOrderNotFound,
		},
	}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			repo := mocks.NewMockRepo(t)
			repo.EXPECT().Exists(ctx, completeInfo.OrderID).Return(tCase.repoExistsResp, nil)
			if !tCase.repoExistsResp {
				repo.EXPECT().CompleteOrder(ctx, completeInfo).Return(tCase.repoCompleteResp, tCase.repoErr)
			}
			useCase := order_usecase.New(repo)
			id, err := useCase.CompleteOrder(ctx, completeInfo)
			require.ErrorIs(t, err, tCase.expErr)
			require.Equal(t, tCase.expResponse, id)
		})
	}
}
