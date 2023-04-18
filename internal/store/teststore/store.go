package teststore

import (
	"github.com/bemmanue/wildberries_L0/internal/model"
	"github.com/bemmanue/wildberries_L0/internal/store"
)

// Store ...
type Store struct {
	orderRepository *OrderRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// Order ...
func (s *Store) Order() store.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		store:  s,
		orders: make(map[string]*model.OrderJSON),
	}

	return s.orderRepository
}
