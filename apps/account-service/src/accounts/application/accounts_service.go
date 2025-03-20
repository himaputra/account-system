package application

import (
	commandhandler "account-system/apps/account-service/src/accounts/application/command-handlers"
	command "account-system/apps/account-service/src/accounts/application/commands"
	query "account-system/apps/account-service/src/accounts/application/queries"
	queryhandler "account-system/apps/account-service/src/accounts/application/query-handlers"
	"account-system/apps/account-service/src/accounts/domain"
)

type AccountsService struct {
	createAccountCommandHandler   *commandhandler.CreateAccountCommandHandler
	depositAccountCommandHandler  *commandhandler.DepositAccountCommandHandler
	withdrawAccountCommandHandler *commandhandler.WithdrawAccountCommandHandler
	getAccountQueryHandler        *queryhandler.GetAccountQueryHandler
}

func NewAccountsService(
	createAccountCommandHandler *commandhandler.CreateAccountCommandHandler,
	depositAccountCommandHandler *commandhandler.DepositAccountCommandHandler,
	withdrawAccountCommandHandler *commandhandler.WithdrawAccountCommandHandler,
	getAccountQueryHandler *queryhandler.GetAccountQueryHandler,
) *AccountsService {
	return &AccountsService{
		createAccountCommandHandler:   createAccountCommandHandler,
		depositAccountCommandHandler:  depositAccountCommandHandler,
		withdrawAccountCommandHandler: withdrawAccountCommandHandler,
		getAccountQueryHandler:        getAccountQueryHandler,
	}
}

func (s *AccountsService) Create(createAccountCommand command.CreateAccountCommand) (*domain.Account, error) {
	return s.createAccountCommandHandler.Execute(createAccountCommand)
}

func (s *AccountsService) Deposit(depositAccountCommand command.DepositAccountCommand) (*domain.Account, error) {
	return s.depositAccountCommandHandler.Execute(depositAccountCommand)
}

func (s *AccountsService) Withdraw(withdrawAccountCommand command.WithdrawAccountCommand) (*domain.Account, error) {
	return s.withdrawAccountCommandHandler.Execute(withdrawAccountCommand)
}

func (s *AccountsService) FindOne(getAccountQuery query.GetAccountQuery) (*domain.Account, error) {
	return s.getAccountQueryHandler.Execute(getAccountQuery)
}
