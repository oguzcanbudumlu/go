package ddd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// OrderFactory is responsible for creating valid Order aggregates
type OrderFactory struct {
}

func (f OrderFactory) NewOrder(orderID, customer string, items []OrderItem) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.New("order must contain at least one item")
	}

	// Create the order with a default status
	order := &Order{
		ID:        orderID,
		Customer:  customer,
		Items:     items,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	return order, nil
}

func TestNewOrder_Success(t *testing.T) {
	// Given a valid order
	price, _ := NewMoney(20.0, "USD")
	item, _ := NewOrderItem("item1", "Laptop", 1, price)
	items := []OrderItem{item}

	orderFactory := OrderFactory{}
	order, err := orderFactory.NewOrder("order123", "John Doe", items)

	// Then the order should be created successfully
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, "order123", order.ID)
	assert.Equal(t, "pending", order.Status)
	assert.Len(t, order.Items, 1)
}

func TestNewOrder_FailsWithNoItems(t *testing.T) {
	orderFactory := OrderFactory{}

	// When creating an order with no items
	order, err := orderFactory.NewOrder("order123", "John Doe", []OrderItem{})

	// Then it should fail
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Equal(t, "order must contain at least one item", err.Error())
}
