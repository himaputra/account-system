package middleware

import (
	exception "account-system/apps/account-service/src/accounts/domain/exceptions"

	"github.com/gofiber/fiber/v2"
)

func AccountExceptionHandlerMiddleware(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}

		if e, ok := err.(*exception.AccountNotFoundException); ok {
			code = fiber.StatusBadRequest
			message = e.Message
		}

		if e, ok := err.(*exception.AccountDomainException); ok {
			code = fiber.StatusBadRequest
			message = e.Message
		}

		return c.Status(code).JSON(fiber.Map{"remark": message})
	}

	return nil
}
