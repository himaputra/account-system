package port

import "account-system/apps/account-service/src/users/domain"

type CreateUserRepository interface {
	SaveUser(user domain.User) (*domain.User, error)
}
