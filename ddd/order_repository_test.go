package ddd

// OrderRepository defines methods for persistence
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id string) (*Order, error)
	FindAll() ([]*Order, error)
}
