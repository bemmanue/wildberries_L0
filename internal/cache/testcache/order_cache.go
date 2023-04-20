package testcache

import (
	"errors"
	"github.com/bemmanue/wildberries_L0/internal/model"
	"github.com/bemmanue/wildberries_L0/internal/store"
)

// OrderCache ...
type OrderCache struct {
	cache  *Cache
	orders map[string]*model.OrderJSON
}

// NewOrderCache ...
func NewOrderCache(store store.Store) (*OrderCache, error) {
	orders, err := store.Order().FindAll()
	if err != nil {
		return nil, err
	}

	c := &OrderCache{
		orders: orders,
	}

	return c, nil
}

// Load ...
func (c *OrderCache) Load(orders map[string]*model.OrderJSON) {
	c.orders = orders
}

// Create ...
func (c *OrderCache) Create(order *model.OrderJSON) error {
	if _, ok := c.orders[order.OrderUID]; ok == true {
		return errors.New("already exists")
	}

	c.orders[order.OrderUID] = order
	return nil
}

// Find ...
func (c *OrderCache) Find(orderUID string) (*model.OrderJSON, error) {
	if order, ok := c.orders[orderUID]; ok == true {
		return order, nil
	}

	return nil, errors.New("record not found")
}
