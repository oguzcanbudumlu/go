package ddd

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Money represents a value object for monetary values
type Money struct {
	amount   float64
	currency string
}

// NewMoney is a constructor enforcing invariants
func NewMoney(amount float64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("amount cannot be negative")
	}
	if currency == "" {
		return Money{}, errors.New("currency cannot be empty")
	}
	return Money{amount: amount, currency: currency}, nil
}

func (m Money) Amount() float64 {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, errors.New("currency mismatch")
	}
	return Money{amount: m.amount + other.amount, currency: m.currency}, nil
}

// Equals checks if two Money objects are equal
func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

func (m Money) String() string {
	return fmt.Sprintf("%.2f %s", m.amount, m.currency)
}

func TestNewMoney_Success(t *testing.T) {
	money, err := NewMoney(100.0, "USD")

	assert.NoError(t, err)
	assert.Equal(t, 100.0, money.Amount())
	assert.Equal(t, "USD", money.Currency())
}

func TestNewMoney_FailsWithNegativeAmount(t *testing.T) {
	_, err := NewMoney(-50.0, "USD")

	assert.Error(t, err)
	assert.Equal(t, "amount cannot be negative", err.Error())
}

func TestMoney_Addition_Success(t *testing.T) {
	money1, _ := NewMoney(50.0, "USD")
	money2, _ := NewMoney(30.0, "USD")

	result, err := money1.Add(money2)

	assert.NoError(t, err)
	assert.Equal(t, 80.0, result.Amount())
}

func TestMoney_Addition_FailsWithCurrencyMismatch(t *testing.T) {
	money1, _ := NewMoney(50.0, "USD")
	money2, _ := NewMoney(30.0, "EUR")

	_, err := money1.Add(money2)

	assert.Error(t, err)
	assert.Equal(t, "currency mismatch", err.Error())
}

func TestMoney_Equality(t *testing.T) {
	money1, _ := NewMoney(100.0, "USD")
	money2, _ := NewMoney(100.0, "USD")
	money3, _ := NewMoney(50.0, "USD")

	assert.True(t, money1.Equals(money2))
	assert.False(t, money1.Equals(money3))
}
