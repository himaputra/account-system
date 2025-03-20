package port

import "account-system/apps/account-service/src/accounts/domain"

type FindAccountRepository interface {
	FindAccount(id string) (*domain.Account, error)
}
