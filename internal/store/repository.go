package store

// OrderRepository ...
type OrderRepository interface {
	Create(string) error
}
