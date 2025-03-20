package command

import "github.com/shopspring/decimal"

type WithdrawAccountCommand struct {
	AccountId string
	Amount    decimal.Decimal
}

func NewWithdrawAccountCommand(accountId string, amount decimal.Decimal) *WithdrawAccountCommand {
	return &WithdrawAccountCommand{
		AccountId: accountId,
		Amount:    amount,
	}
}
