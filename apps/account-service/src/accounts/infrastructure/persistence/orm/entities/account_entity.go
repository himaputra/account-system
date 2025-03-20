package entity

import "github.com/shopspring/decimal"

type AccountEntity struct {
	Id      string `gorm:"primaryKey"`
	UserId  string
	Balance decimal.Decimal `gorm:"type:numeric(20,6);"`
}
