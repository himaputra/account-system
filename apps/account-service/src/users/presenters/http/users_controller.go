package controller

import (
	accounts "account-system/apps/account-service/src/accounts/application"
	accountcommand "account-system/apps/account-service/src/accounts/application/commands"
	"account-system/apps/account-service/src/users/application"
	command "account-system/apps/account-service/src/users/application/commands"
	"account-system/apps/account-service/src/users/presenters/http/dto"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	usersService    *application.UsersService
	accountsService *accounts.AccountsService
}

func NewUsersController(usersService *application.UsersService, accountsService *accounts.AccountsService) *UsersController {
	return &UsersController{
		usersService:    usersService,
		accountsService: accountsService,
	}
}

func (uc *UsersController) Create(c *fiber.Ctx) error {
	body := new(dto.CreateUserDto)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	user, err := uc.usersService.Create(*command.NewCreateUserCommand(
		body.Name,
		body.Nik,
		body.PhoneNumber,
	))
	if err != nil {
		if strings.Contains(err.Error(), "conflict") {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return err
	}

	account, err := uc.accountsService.Create(*accountcommand.NewCreateAccountCommand(user.Id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"no_rekening": account.Id})
}
