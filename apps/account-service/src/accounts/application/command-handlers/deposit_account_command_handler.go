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

type DepositAccountCommandHandler struct {
	logger                  zerolog.Logger
	findAccountRepository   port.FindAccountRepository
	createAccountRepository port.CreateAccountRepository
}

func NewDepositAccountCommandHandler(
	findAcccountRepository port.FindAccountRepository,
	createAccountRepository port.CreateAccountRepository,
) *DepositAccountCommandHandler {
	return &DepositAccountCommandHandler{
		logger:                  common.NewLogger("DepositAccountCommandHandler"),
		findAccountRepository:   findAcccountRepository,
		createAccountRepository: createAccountRepository,
	}
}

func (ch *DepositAccountCommandHandler) Execute(depositAccountCommand command.DepositAccountCommand) (*domain.Account, error) {
	commandJson, _ := json.Marshal(depositAccountCommand)
	ch.logger.Debug().Msg("Processing 'DepositAccountCommand': " + string(commandJson))

	account, err := ch.findAccountRepository.FindAccount(depositAccountCommand.AccountId)
	if err != nil {
		ch.logger.Error().Msg("Error finding account: " + depositAccountCommand.AccountId)
		return nil, err
	}
	if account == nil {
		ch.logger.Error().Msg("Account not found: " + depositAccountCommand.AccountId)
		return nil, exception.NewAccountNotFoundException("Account with id: " + depositAccountCommand.AccountId + " is not found")
	}

	account.Deposit(*valueobjects.NewMoney(depositAccountCommand.Amount.InexactFloat64()))

	updatedAccount, err := ch.createAccountRepository.SaveAccount(*account)
	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}
