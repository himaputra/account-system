package mapper

import (
	"account-system/apps/account-service/src/accounts/domain"
	entity "account-system/apps/account-service/src/accounts/infrastructure/persistence/orm/entities"
	"account-system/apps/common/valueobjects"
)

func ToDomain(accountEntity entity.AccountEntity) *domain.Account {
	model := domain.NewAccount(accountEntity.Id)
	model.UserId = accountEntity.UserId
	model.Balance = *valueobjects.NewMoney(accountEntity.Balance.InexactFloat64())

	return model
}

func ToPersistence(account domain.Account) *entity.AccountEntity {
	return &entity.AccountEntity{
		Id:      account.Id,
		UserId:  account.UserId,
		Balance: account.Balance.GetAmount(),
	}
}
