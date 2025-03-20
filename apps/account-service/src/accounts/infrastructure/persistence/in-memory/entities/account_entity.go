package entity

import "github.com/shopspring/decimal"

type AccountEntity struct {
	Id      string
	UserId  string
	Balance decimal.Decimal
}
