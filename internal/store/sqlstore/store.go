package sqlstore

import (
	"database/sql"
	"github.com/bemmanue/wildberries_L0/internal/store"
)

// Store ...
type Store struct {
	db              *sql.DB
	orderRepository *OrderRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Order ...
func (s *Store) Order() store.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		store: s,
	}

	return s.orderRepository
}
