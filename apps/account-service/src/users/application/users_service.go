package application

import (
	commandhandler "account-system/apps/account-service/src/users/application/command-handlers"
	command "account-system/apps/account-service/src/users/application/commands"
	"account-system/apps/account-service/src/users/domain"
)

type UsersService struct {
	createUserCommandHandler *commandhandler.CreateUserCommandHandler
}

func NewUsersService(createUserCommandHandler *commandhandler.CreateUserCommandHandler) *UsersService {
	return &UsersService{
		createUserCommandHandler: createUserCommandHandler,
	}
}

func (s *UsersService) Create(createUserCommand command.CreateUserCommand) (*domain.User, error) {
	return s.createUserCommandHandler.Execute(createUserCommand)
}
