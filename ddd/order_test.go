package ddd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Order represents the aggregate root
type Order struct {
	ID        string
	Customer  string
	Items     []OrderItem
	Status    string
	CreatedAt time.Time
}

func NewOrder(id, customer string) Order {
	return Order{ID: id, Customer: customer, Items: []OrderItem{}, Status: "pending", CreatedAt: time.Now()}
}

func (o *Order) AddItem(item OrderItem) {
	o.Items = append(o.Items, item)
}

func (o Order) TotalAmount() Money {
	total, _ := NewMoney(0, "USD")

	for _, item := range o.Items {
		total, _ = total.Add(item.TotalPrice())

	}

	return total
}

func (o *Order) Complete() error {
	if len(o.Items) == 0 {
		return errors.New("cannot complete an order with no items")
	}

	o.Status = "completed"

	return nil
}

func TestOrder_Creation(t *testing.T) {
	order := NewOrder("order123", "John Doe")

	assert.Equal(t, "order123", order.ID)
	assert.Equal(t, "John Doe", order.Customer)
	assert.Equal(t, "pending", order.Status)
	assert.Empty(t, order.Items)
	assert.WithinDuration(t, time.Now(), order.CreatedAt, time.Second)
}

func TestOrder_AddItem(t *testing.T) {
	order := NewOrder("order123", "John Doe")
	price, _ := NewMoney(10.0, "USD")
	item, _ := NewOrderItem("item1", "Laptop", 2, price)

	order.AddItem(item)

	assert.Len(t, order.Items, 1)
	assert.Equal(t, "Laptop", order.Items[0].Product)
}

func TestOrder_TotalAmount(t *testing.T) {
	order := NewOrder("order123", "John Doe")
	price, _ := NewMoney(10.0, "USD")
	item1, _ := NewOrderItem("item1", "Laptop", 2, price)
	item2, _ := NewOrderItem("item2", "Mouse", 1, price)

	order.AddItem(item1)
	order.AddItem(item2)

	assert.Equal(t, 30.0, order.TotalAmount().Amount())
}

func TestOrder_Complete(t *testing.T) {
	order := NewOrder("order123", "John Doe")
	price, _ := NewMoney(10.0, "USD")
	item, _ := NewOrderItem("item1", "Laptop", 2, price)
	order.AddItem(item)

	err := order.Complete()

	assert.NoError(t, err)
	assert.Equal(t, "completed", order.Status)
}

func TestOrder_Complete_FailsWithoutItems(t *testing.T) {
	order := NewOrder("order123", "John Doe")

	err := order.Complete()

	assert.Error(t, err)
	assert.Equal(t, "cannot complete an order with no items", err.Error())
}
