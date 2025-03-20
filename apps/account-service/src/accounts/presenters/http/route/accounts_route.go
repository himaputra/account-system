package route

import (
	controller "account-system/apps/account-service/src/accounts/presenters/http"

	"github.com/gofiber/fiber/v2"
)

func AccountsRoute(app *fiber.App, controller *controller.AccountsController) {
	app.Get("/saldo/:accountId", controller.FindOne)
	app.Post("/tabung", controller.Deposit)
	app.Post("/tarik", controller.Withdraw)
}
