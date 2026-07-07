package repository

import "tugas3/entity"

type UserRepository interface{
	Register(user entity.User) (entity.User, error)
	Login(identifier string) (entity.User, error)
	GetUser(user entity.User) (entity.User, error)
}