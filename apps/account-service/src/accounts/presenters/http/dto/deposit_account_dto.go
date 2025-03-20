package dto

import "github.com/shopspring/decimal"

type DepositAccountDto struct {
	AccountId string          `json:"no_rekening"`
	Amount    decimal.Decimal `json:"nominal"`
}
