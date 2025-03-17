package ddd

import "errors"

// OrderItem represents an entity within an order
type OrderItem struct {
	ID       string
	Product  string
	Quantity int
	Price    Money
}

// NewOrderItem creates a new order item
func NewOrderItem(id, product string, quantity int, price Money) (OrderItem, error) {
	if quantity <= 0 {
		return OrderItem{}, errors.New("quantity must be greater than zero")
	}
	return OrderItem{ID: id, Product: product, Quantity: quantity, Price: price}, nil
}

// TotalPrice calculates the total price of the item
func (oi OrderItem) TotalPrice() Money {
	totalAmount := oi.Price.Amount() * float64(oi.Quantity)
	totalPrice, _ := NewMoney(totalAmount, oi.Price.Currency())
	return totalPrice
}
