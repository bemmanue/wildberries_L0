package teststore

import (
	"database/sql"
	"github.com/bemmanue/wildberries_L0/internal/store"
)

// Store ...
type Store struct {
	orderRepository *OrderRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{}
}

// Order ...
func (s *Store) Order() store.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		store:  s,
		orders: make(map[int]string),
	}

	return s.orderRepository
}
