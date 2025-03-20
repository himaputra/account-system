package orm

import (
	"account-system/apps/account-service/src/accounts/domain"
	entity "account-system/apps/account-service/src/accounts/infrastructure/persistence/orm/entities"
	mapper "account-system/apps/account-service/src/accounts/infrastructure/persistence/orm/mappers"
	"errors"

	"gorm.io/gorm"
)

type OrmFindAccountRepository struct {
	db *gorm.DB
}

func NewOrmFindAccountRepository(db *gorm.DB) *OrmFindAccountRepository {
	return &OrmFindAccountRepository{
		db: db,
	}
}

func (r *OrmFindAccountRepository) FindAccount(id string) (*domain.Account, error) {
	var entity entity.AccountEntity
	result := r.db.First(&entity, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return mapper.ToDomain(entity), nil
}
