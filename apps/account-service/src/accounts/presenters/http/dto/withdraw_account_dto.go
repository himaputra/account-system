package dto

import "github.com/shopspring/decimal"

type WithdrawAccountDto struct {
	AccountId string          `json:"no_rekening"`
	Amount    decimal.Decimal `json:"nominal"`
}
