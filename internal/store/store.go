package store

// Store ...
type Store interface {
	Order() OrderRepository
}
