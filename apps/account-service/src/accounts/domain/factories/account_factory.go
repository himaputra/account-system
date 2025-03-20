package factory

import (
	command "account-system/apps/account-service/src/accounts/application/commands"
	"account-system/apps/account-service/src/accounts/domain"
	"account-system/apps/common/valueobjects"

	"github.com/google/uuid"
)

type AccountFactory struct {
}

func NewAccountFactory() *AccountFactory {
	return &AccountFactory{}
}

func (f *AccountFactory) Create(createAccountCommand command.CreateAccountCommand) *domain.Account {
	accountId := uuid.New()
	account := domain.NewAccount(accountId.String())
	account.UserId = createAccountCommand.UserId
	account.Balance = *valueobjects.NewMoney(0)

	return account
}
