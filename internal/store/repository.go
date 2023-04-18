package store

import "github.com/bemmanue/wildberries_L0/internal/model"

// OrderRepository ...
type OrderRepository interface {
	Create(order *model.OrderJSON) error
	FindAll() (map[string]*model.OrderJSON, error)
}
