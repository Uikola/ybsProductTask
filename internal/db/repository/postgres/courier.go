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
	"github.com/lib/pq"
)

type CourierRepository struct {
	db *sql.DB
}

func NewCourierRepository(db *sql.DB) *CourierRepository {
	return &CourierRepository{db: db}
}

func (cr *CourierRepository) CreateCouriers(ctx context.Context, couriers []entity.Courier) error {
	const op = "courierRepository.CreateCouriers"

	query := `INSERT INTO couriers(type, regions, working_hours) VALUES `
	var values []interface{}
	for i, courier := range couriers {
		n := 3 * i

		workingHours, err := json.Marshal(courier.WorkingHours)
		if err != nil {
			return fmt.Errorf("%s: failed to marhal working hours: %w", op, err)
		}

		query += fmt.Sprintf("($%d, $%d, $%d),", n+1, n+2, n+3)
		values = append(values, courier.Type, pq.Array(courier.Regions), workingHours)
	}
	query = query[:len(query)-1]

	_, err := cr.db.ExecContext(ctx, query, values...)
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	return nil
}

func (cr *CourierRepository) GetCourier(ctx context.Context, courierID int) (entity.Courier, error) {
	const op = "courierRepository.GetCourier"

	query := `
	SELECT id, type, regions, working_hours
	FROM couriers
	WHERE id = $1`

	row := cr.db.QueryRowContext(ctx, query, courierID)
	var id int
	var Type entity.CourierType
	var regionsArray pq.Int64Array
	var wHBytes []byte
	var workingHours []types.Interval
	err := row.Scan(&id, &Type, &regionsArray, &wHBytes)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Courier{}, fmt.Errorf("%s:%w", op, errorz.ErrCourierNotFound)
		}
		return entity.Courier{}, fmt.Errorf("%s:%w", op, err)
	}

	regions := make([]int, len(regionsArray))
	for i, el := range regionsArray {
		regions[i] = int(el)
	}

	err = json.Unmarshal(wHBytes, &workingHours)
	if err != nil {
		return entity.Courier{}, fmt.Errorf("%s: failed to unmarshal working hours: %w", op, err)
	}

	return entity.Courier{
		ID:           id,
		Type:         Type,
		Regions:      regions,
		WorkingHours: workingHours,
	}, nil
}

func (cr *CourierRepository) GetCouriers(ctx context.Context, offset, limit int) ([]entity.Courier, error) {
	const op = "courierRepository.GetCouriers"

	query := `
	SELECT id, type, regions, working_hours
	FROM couriers
	LIMIT $1
	OFFSET $2`

	rows, err := cr.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s:%w", op, errorz.ErrNoCouriers)
		}
		return nil, fmt.Errorf("%s:%w", op, err)
	}
	defer rows.Close()

	var id int
	var Type entity.CourierType
	var regionsArray pq.Int64Array
	var workingHours []types.Interval
	var couriers []entity.Courier
	var wHBytes []byte

	for rows.Next() {
		err = rows.Scan(&id, &Type, &regionsArray, &wHBytes)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}

		regions := make([]int, len(regionsArray))
		for i, el := range regionsArray {
			regions[i] = int(el)
		}

		err = json.Unmarshal(wHBytes, &workingHours)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}

		courier := entity.Courier{
			ID:           id,
			Type:         Type,
			Regions:      regions,
			WorkingHours: workingHours,
		}
		couriers = append(couriers, courier)
	}

	return couriers, nil
}
