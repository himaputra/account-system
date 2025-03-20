package mapper

import (
	"account-system/apps/account-service/src/users/domain"
	entity "account-system/apps/account-service/src/users/infrastructure/persistence/orm/entities"
)

func ToDomain(userEntity entity.UserEntity) *domain.User {
	model := domain.NewUser(userEntity.Id)
	model.Name = userEntity.Name
	model.Nik = userEntity.Nik
	model.PhoneNumber = userEntity.PhoneNumber

	return model
}

func ToPersistence(user domain.User) *entity.UserEntity {
	return &entity.UserEntity{
		Id:          user.Id,
		Name:        user.Name,
		Nik:         user.Nik,
		PhoneNumber: user.PhoneNumber,
	}
}
