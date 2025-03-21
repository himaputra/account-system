// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	application2 "account-system/apps/account-service/src/accounts/application"
	commandhandler2 "account-system/apps/account-service/src/accounts/application/command-handlers"
	port2 "account-system/apps/account-service/src/accounts/application/ports"
	"account-system/apps/account-service/src/accounts/application/query-handlers"
	factory2 "account-system/apps/account-service/src/accounts/domain/factories"
	inmemory2 "account-system/apps/account-service/src/accounts/infrastructure/persistence/in-memory/repositories"
	orm2 "account-system/apps/account-service/src/accounts/infrastructure/persistence/orm/repositories"
	controller2 "account-system/apps/account-service/src/accounts/presenters/http"
	"account-system/apps/account-service/src/core"
	"account-system/apps/account-service/src/users/application"
	"account-system/apps/account-service/src/users/application/command-handlers"
	"account-system/apps/account-service/src/users/application/ports"
	"account-system/apps/account-service/src/users/domain/factories"
	"account-system/apps/account-service/src/users/infrastructure/persistence/in-memory/repositories"
	"account-system/apps/account-service/src/users/infrastructure/persistence/orm/repositories"
	"account-system/apps/account-service/src/users/presenters/http"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitInMemoryApp() (*fiber.App, error) {
	inMemoryUserRepository := inmemory.NewInMemoryUserRepository()
	userFactory := factory.NewUserFactory()
	createUserCommandHandler := commandhandler.NewCreateUserCommandHandler(inMemoryUserRepository, userFactory)
	usersService := application.NewUsersService(createUserCommandHandler)
	accountFactory := factory2.NewAccountFactory()
	inMemoryAccountRepository := inmemory2.NewInMemoryAccountRepository()
	createAccountCommandHandler := commandhandler2.NewCreateAccountCommandHandler(accountFactory, inMemoryAccountRepository)
	depositAccountCommandHandler := commandhandler2.NewDepositAccountCommandHandler(inMemoryAccountRepository, inMemoryAccountRepository)
	withdrawAccountCommandHandler := commandhandler2.NewWithdrawAccountCommandHandler(inMemoryAccountRepository, inMemoryAccountRepository)
	getAccountQueryHandler := queryhandler.NewGetAccountQueryHandler(inMemoryAccountRepository)
	accountsService := application2.NewAccountsService(createAccountCommandHandler, depositAccountCommandHandler, withdrawAccountCommandHandler, getAccountQueryHandler)
	usersController := controller.NewUsersController(usersService, accountsService)
	accountsController := controller2.NewAccountsController(accountsService)
	app, err := core.NewFiber(usersController, accountsController)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func InitOrmApp() (*fiber.App, error) {
	db, err := core.NewGorm()
	if err != nil {
		return nil, err
	}
	ormCreateUserRepository := orm.NewOrmUserRepository(db)
	userFactory := factory.NewUserFactory()
	createUserCommandHandler := commandhandler.NewCreateUserCommandHandler(ormCreateUserRepository, userFactory)
	usersService := application.NewUsersService(createUserCommandHandler)
	accountFactory := factory2.NewAccountFactory()
	ormCreateAccountRepository := orm2.NewOrmCreateAccountRepository(db)
	createAccountCommandHandler := commandhandler2.NewCreateAccountCommandHandler(accountFactory, ormCreateAccountRepository)
	ormFindAccountRepository := orm2.NewOrmFindAccountRepository(db)
	depositAccountCommandHandler := commandhandler2.NewDepositAccountCommandHandler(ormFindAccountRepository, ormCreateAccountRepository)
	withdrawAccountCommandHandler := commandhandler2.NewWithdrawAccountCommandHandler(ormFindAccountRepository, ormCreateAccountRepository)
	getAccountQueryHandler := queryhandler.NewGetAccountQueryHandler(ormFindAccountRepository)
	accountsService := application2.NewAccountsService(createAccountCommandHandler, depositAccountCommandHandler, withdrawAccountCommandHandler, getAccountQueryHandler)
	usersController := controller.NewUsersController(usersService, accountsService)
	accountsController := controller2.NewAccountsController(accountsService)
	app, err := core.NewFiber(usersController, accountsController)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// wire.go:

var InMemoryPersistenceSet = wire.NewSet(inmemory.NewInMemoryUserRepository, wire.Bind(new(port.CreateUserRepository), new(*inmemory.InMemoryUserRepository)), inmemory2.NewInMemoryAccountRepository, wire.Bind(new(port2.CreateAccountRepository), new(*inmemory2.InMemoryAccountRepository)), wire.Bind(new(port2.FindAccountRepository), new(*inmemory2.InMemoryAccountRepository)))

var OrmPersistenceSet = wire.NewSet(core.NewGorm, orm.NewOrmUserRepository, wire.Bind(new(port.CreateUserRepository), new(*orm.OrmCreateUserRepository)), orm2.NewOrmCreateAccountRepository, orm2.NewOrmFindAccountRepository, wire.Bind(new(port2.CreateAccountRepository), new(*orm2.OrmCreateAccountRepository)), wire.Bind(new(port2.FindAccountRepository), new(*orm2.OrmFindAccountRepository)))
