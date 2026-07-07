package repository

import "tugas3/entity"

type PublisherRepository interface{
	Add(publisher entity.Publisher) (entity.Publisher, error)
	Update(publisher entity.Publisher) (entity.Publisher, error)
	Delete(id int32) (error)
	GetData(id int32) (entity.Publisher, error)
	GetAllData() ([]entity.Publisher, error)
}