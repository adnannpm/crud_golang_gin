package repository

import "tugas3/entity"

type CategoryRepository interface{
	Add(categori entity.Categori) (entity.Categori, error)
	Update(categori entity.Categori) (entity.Categori, error)
	Delete(id int32) (error)
	GetAllData() ([]entity.Categori, error)
}