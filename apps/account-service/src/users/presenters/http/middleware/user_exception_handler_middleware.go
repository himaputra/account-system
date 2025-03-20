package middleware

import "github.com/gofiber/fiber/v2"

func UserExceptionHandlerMiddleware(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}

		return c.Status(code).JSON(fiber.Map{"remark": message})
	}

	return nil
}
