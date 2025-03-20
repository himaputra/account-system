package inmemory

import (
	"account-system/apps/account-service/src/accounts/domain"
	entity "account-system/apps/account-service/src/accounts/infrastructure/persistence/in-memory/entities"
	mapper "account-system/apps/account-service/src/accounts/infrastructure/persistence/in-memory/mappers"
)

type InMemoryAccountRepository struct {
	accounts map[string]*entity.AccountEntity
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		accounts: make(map[string]*entity.AccountEntity),
	}
}

func (r InMemoryAccountRepository) SaveAccount(account domain.Account) (*domain.Account, error) {
	persistenceModel := mapper.ToPersistence(account)
	r.accounts[persistenceModel.Id] = persistenceModel

	newEntity := r.accounts[persistenceModel.Id]

	return mapper.ToDomain(*newEntity), nil
}

func (r InMemoryAccountRepository) FindAccount(id string) (*domain.Account, error) {
	entity := r.accounts[id]
	if entity != nil {
		return mapper.ToDomain(*entity), nil
	}

	return nil, nil
}
