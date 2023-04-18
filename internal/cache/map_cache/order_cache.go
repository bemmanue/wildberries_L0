package map_cache

import (
	"errors"
	"github.com/bemmanue/wildberries_L0/internal/model"
)

// OrderCache ...
type OrderCache struct {
	cache  *Cache
	orders map[string]*model.OrderJSON
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
