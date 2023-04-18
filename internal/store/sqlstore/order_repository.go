package sqlstore

import (
	"github.com/bemmanue/wildberries_L0/internal/model"
)

// OrderRepository ...
type OrderRepository struct {
	store *Store
}

// Create ...
func (r *OrderRepository) Create(order *model.OrderJSON) error {
	if err := r.store.db.QueryRow(
		"insert into orders (order_uid, data) values ($1, $2)",
		order.OrderUID,
		order.Data,
	); err != nil {
		return err.Err()
	}
	return nil
}

// FindAll ...
func (r *OrderRepository) FindAll() (map[string]*model.OrderJSON, error) {
	orders := make(map[string]*model.OrderJSON)

	rows, err := r.store.db.Query("select order_uid, data from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		order := model.OrderJSON{}
		if err := rows.Scan(
			&order.OrderUID,
			&order.Data,
		); err != nil {
			return nil, err
		}

		orders[order.OrderUID] = &order
	}

	return orders, nil
}
