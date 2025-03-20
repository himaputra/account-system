package commandhandler

import (
	command "account-system/apps/account-service/src/accounts/application/commands"
	port "account-system/apps/account-service/src/accounts/application/ports"
	"account-system/apps/account-service/src/accounts/domain"
	factory "account-system/apps/account-service/src/accounts/domain/factories"
	"account-system/apps/common"
	"encoding/json"

	"github.com/rs/zerolog"
)

type CreateAccountCommandHandler struct {
	logger                  zerolog.Logger
	accountFactory          *factory.AccountFactory
	createAccountRepository port.CreateAccountRepository
}

func NewCreateAccountCommandHandler(
	accountFactory *factory.AccountFactory,
	createAccountRepository port.CreateAccountRepository,
) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		logger:                  common.NewLogger("CreateAccountCommandHandler"),
		accountFactory:          accountFactory,
		createAccountRepository: createAccountRepository,
	}
}

func (ch *CreateAccountCommandHandler) Execute(createAccountCommand command.CreateAccountCommand) (*domain.Account, error) {
	commandJson, _ := json.Marshal(createAccountCommand)
	ch.logger.Debug().Msg("Processing 'CreateAccountCommand': " + string(commandJson))

	account := ch.accountFactory.Create(createAccountCommand)

	return ch.createAccountRepository.SaveAccount(*account)
}
