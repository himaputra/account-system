package factory

import (
	command "account-system/apps/account-service/src/users/application/commands"
	"account-system/apps/account-service/src/users/domain"

	"github.com/google/uuid"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (f *UserFactory) Create(createUserCommand command.CreateUserCommand) *domain.User {
	userId := uuid.New()
	user := domain.NewUser(userId.String())
	user.Name = createUserCommand.Name
	user.Nik = createUserCommand.Nik
	user.PhoneNumber = createUserCommand.PhoneNumber

	return user
}
