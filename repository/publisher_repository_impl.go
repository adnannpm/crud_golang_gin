package repository

import (
	"tugas3/entity"

	"gorm.io/gorm"
)

type PublisherRepositoryImpl struct{
	DB *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) PublisherRepository{
	return &PublisherRepositoryImpl{ DB: db }
}

func (repository *PublisherRepositoryImpl) Add (publisher entity.Publisher) (entity.Publisher, error) {
	result := repository.DB.Create(&publisher)
	if result.Error != nil {
		return publisher, result.Error
	}

	return publisher, nil
}

func (repository *PublisherRepositoryImpl) Update (publisher entity.Publisher) (entity.Publisher, error){
	result := repository.DB.Model(&entity.Publisher{}).Where("id = ?", publisher.ID).Updates(publisher)
	if result.Error != nil {
		return publisher, result.Error
	}
	return publisher, nil
}

func (repository *PublisherRepositoryImpl) Delete (id int32) (error){
	var publisher entity.Publisher
	result := repository.DB.Delete(&publisher, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *PublisherRepositoryImpl) GetData(id int32) (entity.Publisher, error){
	var publisher entity.Publisher
	result := repository.DB.Take(&publisher, "id = ?", id)
	if result.Error != nil {
		return publisher, result.Error
	}
	return publisher, nil
}

func (repository *PublisherRepositoryImpl) GetAllData() ([]entity.Publisher, error){
	var publishers []entity.Publisher
	result := repository.DB.Find(&publishers)
	if result.Error != nil {
		return publishers, result.Error
	}

	return publishers, nil
}