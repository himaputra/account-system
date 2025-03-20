//go:build wireinject
// +build wireinject

package main

import (
	applicationaccounts "account-system/apps/account-service/src/accounts/application"
	commandhandleraccounts "account-system/apps/account-service/src/accounts/application/command-handlers"
	port "account-system/apps/account-service/src/accounts/application/ports"
	portaccounts "account-system/apps/account-service/src/accounts/application/ports"
	queryhandleraccounts "account-system/apps/account-service/src/accounts/application/query-handlers"
	factoryaccounts "account-system/apps/account-service/src/accounts/domain/factories"
	inmemoryaccounts "account-system/apps/account-service/src/accounts/infrastructure/persistence/in-memory/repositories"
	ormaccounts "account-system/apps/account-service/src/accounts/infrastructure/persistence/orm/repositories"
	controlleraccounts "account-system/apps/account-service/src/accounts/presenters/http"
	"account-system/apps/account-service/src/core"
	applicationusers "account-system/apps/account-service/src/users/application"
	commandhandlerusers "account-system/apps/account-service/src/users/application/command-handlers"
	portusers "account-system/apps/account-service/src/users/application/ports"
	factoryusers "account-system/apps/account-service/src/users/domain/factories"
	inmemoryusers "account-system/apps/account-service/src/users/infrastructure/persistence/in-memory/repositories"
	ormusers "account-system/apps/account-service/src/users/infrastructure/persistence/orm/repositories"
	controllerusers "account-system/apps/account-service/src/users/presenters/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var InMemoryPersistenceSet = wire.NewSet(
	inmemoryusers.NewInMemoryUserRepository,
	wire.Bind(new(portusers.CreateUserRepository), new(*inmemoryusers.InMemoryUserRepository)),
	inmemoryaccounts.NewInMemoryAccountRepository,
	wire.Bind(new(portaccounts.CreateAccountRepository), new(*inmemoryaccounts.InMemoryAccountRepository)),
	wire.Bind(new(portaccounts.FindAccountRepository), new(*inmemoryaccounts.InMemoryAccountRepository)),
)
var OrmPersistenceSet = wire.NewSet(
	core.NewGorm,
	ormusers.NewOrmUserRepository,
	wire.Bind(new(portusers.CreateUserRepository), new(*ormusers.OrmCreateUserRepository)),
	ormaccounts.NewOrmCreateAccountRepository,
	ormaccounts.NewOrmFindAccountRepository,
	wire.Bind(new(port.CreateAccountRepository), new(*ormaccounts.OrmCreateAccountRepository)),
	wire.Bind(new(port.FindAccountRepository), new(*ormaccounts.OrmFindAccountRepository)),
)

func InitInMemoryApp() (*fiber.App, error) {
	wire.Build(
		core.NewFiber,
		controllerusers.NewUsersController,
		applicationusers.NewUsersService,
		commandhandlerusers.NewCreateUserCommandHandler,
		factoryusers.NewUserFactory,
		InMemoryPersistenceSet,
		controlleraccounts.NewAccountsController,
		applicationaccounts.NewAccountsService,
		factoryaccounts.NewAccountFactory,
		commandhandleraccounts.NewCreateAccountCommandHandler,
		commandhandleraccounts.NewDepositAccountCommandHandler,
		commandhandleraccounts.NewWithdrawAccountCommandHandler,
		queryhandleraccounts.NewGetAccountQueryHandler,
	)

	return &fiber.App{}, nil
}

func InitOrmApp() (*fiber.App, error) {
	wire.Build(
		core.NewFiber,
		controllerusers.NewUsersController,
		applicationusers.NewUsersService,
		commandhandlerusers.NewCreateUserCommandHandler,
		factoryusers.NewUserFactory,
		OrmPersistenceSet,
		controlleraccounts.NewAccountsController,
		applicationaccounts.NewAccountsService,
		factoryaccounts.NewAccountFactory,
		commandhandleraccounts.NewCreateAccountCommandHandler,
		commandhandleraccounts.NewDepositAccountCommandHandler,
		commandhandleraccounts.NewWithdrawAccountCommandHandler,
		queryhandleraccounts.NewGetAccountQueryHandler,
	)

	return &fiber.App{}, nil
}
