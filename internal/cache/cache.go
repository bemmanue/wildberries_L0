package cache

import "github.com/bemmanue/wildberries_L0/internal/model"

// Cache ...
type Cache interface {
	Order() OrderCache
}

// OrderCache ...
type OrderCache interface {
	Load(orders map[string]*model.OrderJSON)
	Create(order *model.OrderJSON) error
	Find(orderUID string) (*model.OrderJSON, error)
}
