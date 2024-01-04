package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Uikola/ybsProductTask/internal/entity"
	"github.com/Uikola/ybsProductTask/internal/entity/types"
	"github.com/Uikola/ybsProductTask/internal/errorz"
	"time"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) CreateOrders(ctx context.Context, orders []entity.Order) error {
	const op = "orderRepository.CreateOrders"

	query := `INSERT INTO orders(weight, region, delivery_time, price) VALUES `
	var values []interface{}
	for i, order := range orders {
		n := 4 * i

		deliveryTime, err := json.Marshal(order.DeliveryTime)
		if err != nil {
			return fmt.Errorf("%s: failed to unmarhsal delivery time: %w", op, err)
		}

		query += fmt.Sprintf("($%d, $%d, $%d, $%d),", n+1, n+2, n+3, n+4)
		values = append(values, order.Weight, order.Region, deliveryTime, order.Price)
	}
	query = query[:len(query)-1]

	_, err := or.db.ExecContext(ctx, query, values...)
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	return nil
}

func (or *OrderRepository) GetOrder(ctx context.Context, orderID int) (entity.Order, error) {
	const op = "orderRepository.GetOrder"

	query := `
	SELECT id, weight, region, delivery_time, price, complete_time, courier_id
	FROM orders
	WHERE id = $1`

	row := or.db.QueryRowContext(ctx, query, orderID)

	var id, weight, region, price int
	var courierID *int
	var dtBytes []byte
	var deliveryTime []types.Interval
	var completeTime *time.Time
	err := row.Scan(&id, &weight, &region, &dtBytes, &price, &completeTime, &courierID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Order{}, fmt.Errorf("%s:%w", op, errorz.ErrOrderNotFound)
		}
		return entity.Order{}, fmt.Errorf("%s:%w", op, err)
	}

	err = json.Unmarshal(dtBytes, &deliveryTime)
	if err != nil {
		return entity.Order{}, fmt.Errorf("%s: failed to unmarshal hours: %w", op, err)
	}

	return entity.Order{
		ID:           id,
		Weight:       weight,
		Region:       region,
		DeliveryTime: deliveryTime,
		Price:        price,
		CompleteTime: completeTime,
		CourierID:    courierID,
	}, nil
}

func (or *OrderRepository) GetOrders(ctx context.Context, offset, limit int) ([]entity.Order, error) {
	const op = "orderRepository.GetOrders"

	query := `
	SELECT id, weight, region, delivery_time, price, complete_time, courier_id
	FROM orders
	LIMIT $1
	OFFSET $2`

	rows, err := or.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s:%w", op, errorz.ErrNoOrders)
		}
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	var id, weight, region, price int
	var courierID *int
	var dtBytes []byte
	var deliveryTime []types.Interval
	var completeTime *time.Time
	var orders []entity.Order

	for rows.Next() {
		err = rows.Scan(&id, &weight, &region, &dtBytes, &price, &completeTime, &courierID)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}

		err = json.Unmarshal(dtBytes, &deliveryTime)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to unmarshal delivery time: %w", op, err)
		}

		order := entity.Order{
			ID:           id,
			Weight:       weight,
			Region:       region,
			DeliveryTime: deliveryTime,
			Price:        price,
			CompleteTime: completeTime,
			CourierID:    courierID,
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (or *OrderRepository) CompleteOrder(ctx context.Context, completeInfo entity.CompleteOrderInfo) (int, error) {
	const op = "orderRepository.CompleteOrder"

	query := `
	UPDATE orders 
	SET courier_id = $1,
	    complete_time = $2
	WHERE id = $3
	RETURNING id`

	row := or.db.QueryRowContext(ctx, query, completeInfo.CourierID, completeInfo.CompleteTime, completeInfo.OrderID)
	var id int
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("%s:%w", op, errorz.ErrOrderNotFound)
		}
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	return id, nil
}

func (or *OrderRepository) Exists(ctx context.Context, orderID int) (bool, error) {
	const op = "orderRepository.Exists"

	query := `
	SELECT courier_id
	FROM orders
	WHERE id = $1`

	var courierID *int
	row := or.db.QueryRowContext(ctx, query, orderID)
	err := row.Scan(&courierID)
	if err != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("%s:%w", op, err)
	}

	if courierID != nil {
		return true, nil
	}

	return false, nil
}

func (or *OrderRepository) GetOrdersByCourier(ctx context.Context, courierID int) ([]entity.Order, error) {
	const op = "orderRepository.GetOrdersByCourier"

	query := `
	SELECT id, price, complete_time, courier_id
	FROM orders
	WHERE courier_id = $1`

	rows, err := or.db.QueryContext(ctx, query, courierID)

	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	var id, price int
	var deliveryTime []types.Interval
	var completeTime *time.Time
	var orders []entity.Order

	for rows.Next() {
		err = rows.Scan(&id, &price, &completeTime, &courierID)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}

		order := entity.Order{
			ID:           id,
			DeliveryTime: deliveryTime,
			Price:        price,
			CompleteTime: completeTime,
			CourierID:    &courierID,
		}
		orders = append(orders, order)
	}

	return orders, nil
}
