package valueobjects

import (
	"github.com/shopspring/decimal"
)

type Money struct {
	amount decimal.Decimal
}

func NewMoney(amount float64) *Money {
	return &Money{
		amount: decimal.NewFromFloat(amount),
	}
}

func (m *Money) IsGreaterThanZero() bool {
	return m.amount.IsPositive()
}

func (m *Money) IsLessThan(money Money) bool {
	return m.amount.LessThan(money.amount)
}

func (m *Money) IsGreaterThan(money Money) bool {
	return m.amount.GreaterThan(money.amount)
}

func (m *Money) Add(money Money) Money {
	return *NewMoney(m.setScale(m.amount.Add(money.GetAmount())).InexactFloat64())
}

func (m *Money) Subtract(money Money) Money {
	return *NewMoney(m.setScale(m.amount.Sub(money.GetAmount())).InexactFloat64())
}

func (m *Money) Multiply(money Money) Money {
	return *NewMoney(m.setScale(m.amount.Mul(money.GetAmount())).InexactFloat64())
}

func (m *Money) GetAmount() decimal.Decimal {
	return m.amount
}

func (m *Money) setScale(input decimal.Decimal) decimal.Decimal {
	return input.RoundBank(2)
}
