package mapcache

import (
	"github.com/bemmanue/wildberries_L0/internal/cache"
	"github.com/bemmanue/wildberries_L0/internal/model"
	"github.com/bemmanue/wildberries_L0/internal/store"
)

type Cache struct {
	orderCache *OrderCache
}

// New ...
func New(store store.Store) (*Cache, error) {
	orderCache, err := NewOrderCache(store)
	if err != nil {
		return nil, err
	}

	c := &Cache{
		orderCache: orderCache,
	}

	return c, nil
}

// Order ...
func (c *Cache) Order() cache.OrderCache {
	if c.orderCache != nil {
		return c.orderCache
	}

	c.orderCache = &OrderCache{
		cache:  c,
		orders: make(map[string]*model.OrderJSON),
	}

	return c.orderCache
}
