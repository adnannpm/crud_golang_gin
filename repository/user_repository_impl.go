package repository

import (
	"tugas3/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct{
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) Register(user entity.User) (entity.User, error) {
	result := repository.DB.Create(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Login(identifier string) (entity.User, error){
	var data entity.User
	result := repository.DB.Where("email = ? OR full_name = ?", identifier, identifier).Find(&data)

	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (repository *UserRepositoryImpl) GetUser(user entity.User) (entity.User, error){
	result := repository.DB.Where("id = ?", user.ID).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}