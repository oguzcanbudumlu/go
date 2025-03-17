package ddd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// AnemicOrder Anemic Model (Business logic in service)
type AnemicOrder struct {
	ID     string
	Amount float64
	Status string
}

func CompleteAnemicOrder(order *AnemicOrder) error {
	if order.Status != "pending" {
		return errors.New("order is not pending")
	}
	order.Status = "completed"
	return nil
}

// RichOrder Rich Model (Business logic in entity)
type RichOrder struct {
	ID     string
	Amount float64
	Status string
}

func (o *RichOrder) Complete() error {
	if o.Status != "pending" {
		return errors.New("order is not pending")
	}
	o.Status = "completed"
	return nil
}

// TestAnemicOrder_Complete Tests
func TestAnemicOrder_Complete(t *testing.T) {
	order := &AnemicOrder{ID: "123", Amount: 100.0, Status: "pending"}

	err := CompleteAnemicOrder(order)

	assert.NoError(t, err)
	assert.Equal(t, "completed", order.Status)
}

func TestRichOrder_Complete(t *testing.T) {
	order := &RichOrder{ID: "123", Amount: 100.0, Status: "pending"}

	err := order.Complete()

	assert.NoError(t, err)
	assert.Equal(t, "completed", order.Status)
}

func TestRichOrder_Complete_Fails_WhenNotPending(t *testing.T) {
	order := &RichOrder{ID: "123", Amount: 100.0, Status: "shipped"}

	err := order.Complete()

	assert.Error(t, err)
	assert.Equal(t, "order is not pending", err.Error())
}
