package commandhandler

import (
	command "account-system/apps/account-service/src/users/application/commands"
	port "account-system/apps/account-service/src/users/application/ports"
	"account-system/apps/account-service/src/users/domain"
	factory "account-system/apps/account-service/src/users/domain/factories"
	"account-system/apps/common"
	"encoding/json"

	"github.com/rs/zerolog"
)

type CreateUserCommandHandler struct {
	logger               zerolog.Logger
	createUserRepository port.CreateUserRepository
	userFactory          *factory.UserFactory
}

func NewCreateUserCommandHandler(createUserRepository port.CreateUserRepository, userFactory *factory.UserFactory) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		logger:               common.NewLogger("CreateUserCommandHandler"),
		createUserRepository: createUserRepository,
		userFactory:          userFactory,
	}
}

func (ch *CreateUserCommandHandler) Execute(createUserCommand command.CreateUserCommand) (*domain.User, error) {
	commandJson, _ := json.Marshal(createUserCommand)
	ch.logger.Debug().Msg("Processing 'CreateUserCommand': " + string(commandJson))

	user := ch.userFactory.Create(createUserCommand)

	user, err := ch.createUserRepository.SaveUser(*user)
	if err != nil {
		ch.logger.Error().Msg("Error while saving user: " + err.Error())
		return nil, err
	}

	return user, nil
}
