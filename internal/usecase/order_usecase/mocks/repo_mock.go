// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/Uikola/ybsProductTask/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockRepo is an autogenerated mock type for the Repo type
type MockRepo struct {
	mock.Mock
}

type MockRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepo) EXPECT() *MockRepo_Expecter {
	return &MockRepo_Expecter{mock: &_m.Mock}
}

// CompleteOrder provides a mock function with given fields: ctx, completeInfo
func (_m *MockRepo) CompleteOrder(ctx context.Context, completeInfo entity.CompleteOrderInfo) (int, error) {
	ret := _m.Called(ctx, completeInfo)

	if len(ret) == 0 {
		panic("no return value specified for CompleteOrder")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CompleteOrderInfo) (int, error)); ok {
		return rf(ctx, completeInfo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.CompleteOrderInfo) int); ok {
		r0 = rf(ctx, completeInfo)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.CompleteOrderInfo) error); ok {
		r1 = rf(ctx, completeInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_CompleteOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteOrder'
type MockRepo_CompleteOrder_Call struct {
	*mock.Call
}

// CompleteOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - completeInfo entity.CompleteOrderInfo
func (_e *MockRepo_Expecter) CompleteOrder(ctx interface{}, completeInfo interface{}) *MockRepo_CompleteOrder_Call {
	return &MockRepo_CompleteOrder_Call{Call: _e.mock.On("CompleteOrder", ctx, completeInfo)}
}

func (_c *MockRepo_CompleteOrder_Call) Run(run func(ctx context.Context, completeInfo entity.CompleteOrderInfo)) *MockRepo_CompleteOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.CompleteOrderInfo))
	})
	return _c
}

func (_c *MockRepo_CompleteOrder_Call) Return(_a0 int, _a1 error) *MockRepo_CompleteOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_CompleteOrder_Call) RunAndReturn(run func(context.Context, entity.CompleteOrderInfo) (int, error)) *MockRepo_CompleteOrder_Call {
	_c.Call.Return(run)
	return _c
}

// CreateOrders provides a mock function with given fields: ctx, orders
func (_m *MockRepo) CreateOrders(ctx context.Context, orders []entity.Order) error {
	ret := _m.Called(ctx, orders)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrders")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []entity.Order) error); ok {
		r0 = rf(ctx, orders)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepo_CreateOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrders'
type MockRepo_CreateOrders_Call struct {
	*mock.Call
}

// CreateOrders is a helper method to define mock.On call
//   - ctx context.Context
//   - orders []entity.Order
func (_e *MockRepo_Expecter) CreateOrders(ctx interface{}, orders interface{}) *MockRepo_CreateOrders_Call {
	return &MockRepo_CreateOrders_Call{Call: _e.mock.On("CreateOrders", ctx, orders)}
}

func (_c *MockRepo_CreateOrders_Call) Run(run func(ctx context.Context, orders []entity.Order)) *MockRepo_CreateOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]entity.Order))
	})
	return _c
}

func (_c *MockRepo_CreateOrders_Call) Return(_a0 error) *MockRepo_CreateOrders_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_CreateOrders_Call) RunAndReturn(run func(context.Context, []entity.Order) error) *MockRepo_CreateOrders_Call {
	_c.Call.Return(run)
	return _c
}

// Exists provides a mock function with given fields: ctx, orderID
func (_m *MockRepo) Exists(ctx context.Context, orderID int) (bool, error) {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (bool, error)); ok {
		return rf(ctx, orderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) bool); ok {
		r0 = rf(ctx, orderID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type MockRepo_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - ctx context.Context
//   - orderID int
func (_e *MockRepo_Expecter) Exists(ctx interface{}, orderID interface{}) *MockRepo_Exists_Call {
	return &MockRepo_Exists_Call{Call: _e.mock.On("Exists", ctx, orderID)}
}

func (_c *MockRepo_Exists_Call) Run(run func(ctx context.Context, orderID int)) *MockRepo_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRepo_Exists_Call) Return(_a0 bool, _a1 error) *MockRepo_Exists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_Exists_Call) RunAndReturn(run func(context.Context, int) (bool, error)) *MockRepo_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrder provides a mock function with given fields: ctx, orderID
func (_m *MockRepo) GetOrder(ctx context.Context, orderID int) (entity.Order, error) {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrder")
	}

	var r0 entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (entity.Order, error)); ok {
		return rf(ctx, orderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) entity.Order); ok {
		r0 = rf(ctx, orderID)
	} else {
		r0 = ret.Get(0).(entity.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrder'
type MockRepo_GetOrder_Call struct {
	*mock.Call
}

// GetOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - orderID int
func (_e *MockRepo_Expecter) GetOrder(ctx interface{}, orderID interface{}) *MockRepo_GetOrder_Call {
	return &MockRepo_GetOrder_Call{Call: _e.mock.On("GetOrder", ctx, orderID)}
}

func (_c *MockRepo_GetOrder_Call) Run(run func(ctx context.Context, orderID int)) *MockRepo_GetOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRepo_GetOrder_Call) Return(_a0 entity.Order, _a1 error) *MockRepo_GetOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetOrder_Call) RunAndReturn(run func(context.Context, int) (entity.Order, error)) *MockRepo_GetOrder_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrders provides a mock function with given fields: ctx, offset, limit
func (_m *MockRepo) GetOrders(ctx context.Context, offset int, limit int) ([]entity.Order, error) {
	ret := _m.Called(ctx, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 []entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]entity.Order, error)); ok {
		return rf(ctx, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []entity.Order); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrders'
type MockRepo_GetOrders_Call struct {
	*mock.Call
}

// GetOrders is a helper method to define mock.On call
//   - ctx context.Context
//   - offset int
//   - limit int
func (_e *MockRepo_Expecter) GetOrders(ctx interface{}, offset interface{}, limit interface{}) *MockRepo_GetOrders_Call {
	return &MockRepo_GetOrders_Call{Call: _e.mock.On("GetOrders", ctx, offset, limit)}
}

func (_c *MockRepo_GetOrders_Call) Run(run func(ctx context.Context, offset int, limit int)) *MockRepo_GetOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockRepo_GetOrders_Call) Return(_a0 []entity.Order, _a1 error) *MockRepo_GetOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetOrders_Call) RunAndReturn(run func(context.Context, int, int) ([]entity.Order, error)) *MockRepo_GetOrders_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrdersByCourier provides a mock function with given fields: ctx, courierID
func (_m *MockRepo) GetOrdersByCourier(ctx context.Context, courierID int) ([]entity.Order, error) {
	ret := _m.Called(ctx, courierID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrdersByCourier")
	}

	var r0 []entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]entity.Order, error)); ok {
		return rf(ctx, courierID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []entity.Order); ok {
		r0 = rf(ctx, courierID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, courierID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetOrdersByCourier_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrdersByCourier'
type MockRepo_GetOrdersByCourier_Call struct {
	*mock.Call
}

// GetOrdersByCourier is a helper method to define mock.On call
//   - ctx context.Context
//   - courierID int
func (_e *MockRepo_Expecter) GetOrdersByCourier(ctx interface{}, courierID interface{}) *MockRepo_GetOrdersByCourier_Call {
	return &MockRepo_GetOrdersByCourier_Call{Call: _e.mock.On("GetOrdersByCourier", ctx, courierID)}
}

func (_c *MockRepo_GetOrdersByCourier_Call) Run(run func(ctx context.Context, courierID int)) *MockRepo_GetOrdersByCourier_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRepo_GetOrdersByCourier_Call) Return(_a0 []entity.Order, _a1 error) *MockRepo_GetOrdersByCourier_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetOrdersByCourier_Call) RunAndReturn(run func(context.Context, int) ([]entity.Order, error)) *MockRepo_GetOrdersByCourier_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepo creates a new instance of MockRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepo {
	mock := &MockRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}