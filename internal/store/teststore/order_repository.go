package teststore

// OrderRepository ...
type OrderRepository struct {
	store  *Store
	orders map[int]string
}

// Create ...
func (r *OrderRepository) Create(string) error {
	return nil
}
