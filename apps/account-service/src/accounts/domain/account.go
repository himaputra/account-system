package domain

import (
	"account-system/apps/common/valueobjects"
)

type Account struct {
	Id      string
	UserId  string
	Balance valueobjects.Money
}

func NewAccount(id string) *Account {
	return &Account{
		Id: id,
	}
}

func (a *Account) Deposit(money valueobjects.Money) {
	a.Balance = a.Balance.Add(money)
}

func (a *Account) Withdraw(money valueobjects.Money) {
	a.Balance = a.Balance.Subtract(money)
}
