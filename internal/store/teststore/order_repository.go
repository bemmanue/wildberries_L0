package teststore

import (
	"errors"
	"github.com/bemmanue/wildberries_L0/internal/model"
)

// OrderRepository ...
type OrderRepository struct {
	store  *Store
	orders map[string]*model.OrderJSON
}

// Create ...
func (r *OrderRepository) Create(order *model.OrderJSON) error {
	if _, ok := r.orders[order.OrderUID]; ok == true {
		return errors.New("record already exists")
	}

	r.orders[order.OrderUID] = order
	return nil
}

// FindAll ...
func (r *OrderRepository) FindAll() (map[string]*model.OrderJSON, error) {
	orders := make(map[string]*model.OrderJSON)

	for uid, order := range orders {
		orders[uid] = order
	}

	return orders, nil
}
