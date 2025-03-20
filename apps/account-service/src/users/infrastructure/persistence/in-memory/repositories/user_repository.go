package inmemory

import (
	"account-system/apps/account-service/src/users/domain"
	entity "account-system/apps/account-service/src/users/infrastructure/persistence/in-memory/entities"
	mapper "account-system/apps/account-service/src/users/infrastructure/persistence/in-memory/mappers"
)

type InMemoryUserRepository struct {
	users map[string]*entity.UserEntity
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entity.UserEntity),
	}
}

func (r *InMemoryUserRepository) SaveUser(user domain.User) (*domain.User, error) {
	persistenceModel := mapper.ToPersistence(user)
	r.users[persistenceModel.Id] = persistenceModel

	newEntity := r.users[persistenceModel.Id]

	return mapper.ToDomain(*newEntity), nil
}
