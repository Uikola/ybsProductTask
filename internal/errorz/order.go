package errorz

import "errors"

var ErrOrderNotFound = errors.New("order not found")
var ErrNoOrders = errors.New("no orders")
var ErrOrderAlreadyExists = errors.New("order already exists")
var ErrInvalidWeight = errors.New("invalid weight")
var ErrInvalidPrice = errors.New("invalid price")
var ErrInvalidDeliveryTime = errors.New("invalid delivery time")
