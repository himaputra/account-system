package controller

import (
	"account-system/apps/account-service/src/accounts/application"
	command "account-system/apps/account-service/src/accounts/application/commands"
	query "account-system/apps/account-service/src/accounts/application/queries"
	"account-system/apps/account-service/src/accounts/presenters/http/dto"

	"github.com/gofiber/fiber/v2"
)

type AccountsController struct {
	accountsService *application.AccountsService
}

func NewAccountsController(accountsService *application.AccountsService) *AccountsController {
	return &AccountsController{
		accountsService: accountsService,
	}
}

func (ac *AccountsController) FindOne(c *fiber.Ctx) error {
	accountId := c.Params("accountId")
	data, err := ac.accountsService.FindOne(*query.NewGetAccountQuery(accountId))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetAccountDto{Amount: data.Balance.GetAmount().InexactFloat64()})
}

func (ac *AccountsController) Deposit(c *fiber.Ctx) error {
	body := new(dto.DepositAccountDto)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	data, err := ac.accountsService.Deposit(*command.NewDepositAccountCommand(body.AccountId, body.Amount))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetAccountDto{Amount: data.Balance.GetAmount().InexactFloat64()})
}

func (ac *AccountsController) Withdraw(c *fiber.Ctx) error {
	body := new(dto.WithdrawAccountDto)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	data, err := ac.accountsService.Withdraw(*command.NewWithdrawAccountCommand(body.AccountId, body.Amount))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetAccountDto{Amount: data.Balance.GetAmount().InexactFloat64()})
}
