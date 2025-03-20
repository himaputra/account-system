package core

import (
	controlleraccounts "account-system/apps/account-service/src/accounts/presenters/http"
	middlewareaccounts "account-system/apps/account-service/src/accounts/presenters/http/middleware"
	routeaccounts "account-system/apps/account-service/src/accounts/presenters/http/route"
	controllerusers "account-system/apps/account-service/src/users/presenters/http"
	middlewareusers "account-system/apps/account-service/src/users/presenters/http/middleware"
	routeusers "account-system/apps/account-service/src/users/presenters/http/route"

	"github.com/gofiber/fiber/v2"
)

func NewFiber(usersController *controllerusers.UsersController, accountsController *controlleraccounts.AccountsController) (*fiber.App, error) {
	app := fiber.New()

	app.Use(middlewareusers.UserExceptionHandlerMiddleware)
	app.Use(middlewareaccounts.AccountExceptionHandlerMiddleware)

	routeusers.UsersRoute(app, usersController)
	routeaccounts.AccountsRoute(app, accountsController)

	return app, nil
}
