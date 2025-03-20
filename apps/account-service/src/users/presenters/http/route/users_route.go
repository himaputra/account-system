package route

import (
	controller "account-system/apps/account-service/src/users/presenters/http"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, controller *controller.UsersController) {
	app.Post("/daftar", controller.Create)
}
