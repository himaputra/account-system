package port

import "account-system/apps/account-service/src/accounts/domain"

type CreateAccountRepository interface {
	SaveAccount(account domain.Account) (*domain.Account, error)
}
