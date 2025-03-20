package commandhandler

import (
	command "account-system/apps/account-service/src/accounts/application/commands"
	port "account-system/apps/account-service/src/accounts/application/ports"
	"account-system/apps/account-service/src/accounts/domain"
	exception "account-system/apps/account-service/src/accounts/domain/exceptions"
	"account-system/apps/common"
	"account-system/apps/common/valueobjects"
	"encoding/json"

	"github.com/rs/zerolog"
)

type WithdrawAccountCommandHandler struct {
	logger                  zerolog.Logger
	findAccountRepository   port.FindAccountRepository
	createAccountRepository port.CreateAccountRepository
}

func NewWithdrawAccountCommandHandler(
	findAccountRepository port.FindAccountRepository,
	createAccountRepository port.CreateAccountRepository,
) *WithdrawAccountCommandHandler {
	return &WithdrawAccountCommandHandler{
		logger:                  common.NewLogger("WithdrawAccountCommandHandler"),
		findAccountRepository:   findAccountRepository,
		createAccountRepository: createAccountRepository,
	}
}

func (ch *WithdrawAccountCommandHandler) Execute(withdrawAccountCommand command.WithdrawAccountCommand) (*domain.Account, error) {
	commandJson, _ := json.Marshal(withdrawAccountCommand)
	ch.logger.Debug().Msg("Processing 'WithdrawAccountCommand': " + string(commandJson))

	account, err := ch.findAccountRepository.FindAccount(withdrawAccountCommand.AccountId)
	if err != nil {
		ch.logger.Error().Msg("Error finding account: " + withdrawAccountCommand.AccountId)
		return nil, err
	}
	if account == nil {
		ch.logger.Error().Msg("Account not found: " + withdrawAccountCommand.AccountId)
		return nil, exception.NewAccountNotFoundException("Account with id: " + withdrawAccountCommand.AccountId + " is not found")
	}

	if account.Balance.IsLessThan(*valueobjects.NewMoney(withdrawAccountCommand.Amount.InexactFloat64())) {
		ch.logger.Error().Msg("Insufficient account balance: " + withdrawAccountCommand.AccountId)
		return nil, exception.NewAccountDomainException("Insufficient account balance: " + withdrawAccountCommand.AccountId)
	}

	account.Withdraw(*valueobjects.NewMoney(withdrawAccountCommand.Amount.InexactFloat64()))

	updatedAccount, err := ch.createAccountRepository.SaveAccount(*account)
	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}
