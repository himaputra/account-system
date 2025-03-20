package orm

import (
	"account-system/apps/account-service/src/accounts/domain"
	mapper "account-system/apps/account-service/src/accounts/infrastructure/persistence/orm/mappers"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrmCreateAccountRepository struct {
	db *gorm.DB
}

func NewOrmCreateAccountRepository(db *gorm.DB) *OrmCreateAccountRepository {
	return &OrmCreateAccountRepository{
		db: db,
	}
}

func (r *OrmCreateAccountRepository) SaveAccount(account domain.Account) (*domain.Account, error) {
	persistenceModel := mapper.ToPersistence(account)
	result := r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(persistenceModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return mapper.ToDomain(*persistenceModel), nil
}
