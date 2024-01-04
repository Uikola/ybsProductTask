package errorz

import "errors"

var ErrCourierNotFound = errors.New("courier not found")
var ErrNoCouriers = errors.New("no couriers")
var ErrInvalidCourierType = errors.New("invalid courier type")
var ErrInvalidRegion = errors.New("invalid region")
var ErrInvalidWorkingHours = errors.New("invalid working hours")
