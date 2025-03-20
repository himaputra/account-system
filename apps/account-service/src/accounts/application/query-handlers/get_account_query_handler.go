package queryhandler

import (
	port "account-system/apps/account-service/src/accounts/application/ports"
	query "account-system/apps/account-service/src/accounts/application/queries"
	"account-system/apps/account-service/src/accounts/domain"
	exception "account-system/apps/account-service/src/accounts/domain/exceptions"
	"account-system/apps/common"

	"github.com/rs/zerolog"
)

type GetAccountQueryHandler struct {
	logger                zerolog.Logger
	findAccountRepository port.FindAccountRepository
}

func NewGetAccountQueryHandler(findAccountRepository port.FindAccountRepository) *GetAccountQueryHandler {
	return &GetAccountQueryHandler{
		logger:                common.NewLogger("GetAccountQueryHandler"),
		findAccountRepository: findAccountRepository,
	}
}

func (qh *GetAccountQueryHandler) Execute(getAccountQuery query.GetAccountQuery) (*domain.Account, error) {
	account, err := qh.findAccountRepository.FindAccount(getAccountQuery.Id)
	if err != nil {
		qh.logger.Error().Msg("Error finding Account: " + getAccountQuery.Id)
		return nil, err
	}
	if account == nil {
		return nil, exception.NewAccountNotFoundException("Account with id: " + getAccountQuery.Id + " is not found")
	}

	return account, nil
}
