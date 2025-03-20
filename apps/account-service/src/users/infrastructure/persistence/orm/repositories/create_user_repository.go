package orm

import (
	"account-system/apps/account-service/src/users/domain"
	entity "account-system/apps/account-service/src/users/infrastructure/persistence/orm/entities"
	mapper "account-system/apps/account-service/src/users/infrastructure/persistence/orm/mappers"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrmCreateUserRepository struct {
	db *gorm.DB
}

func NewOrmUserRepository(db *gorm.DB) *OrmCreateUserRepository {
	return &OrmCreateUserRepository{
		db: db,
	}
}

func (r *OrmCreateUserRepository) SaveUser(user domain.User) (*domain.User, error) {
	var existingUser *entity.UserEntity
	if err := r.db.Where("nik = ? OR phone_number = ?", user.Nik, user.PhoneNumber).First(&existingUser).Error; err == nil {
		// Jika ditemukan konflik pada Nik atau PhoneNumber, return error
		return nil, errors.New("conflict: nik or phone_number already exists")
	}

	persistenceModel := mapper.ToPersistence(user)
	result := r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(persistenceModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return mapper.ToDomain(*persistenceModel), nil
}
