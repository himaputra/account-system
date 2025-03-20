package command

import "github.com/shopspring/decimal"

type DepositAccountCommand struct {
	AccountId string
	Amount    decimal.Decimal
}

func NewDepositAccountCommand(accountId string, amount decimal.Decimal) *DepositAccountCommand {
	return &DepositAccountCommand{
		AccountId: accountId,
		Amount:    amount,
	}
}
