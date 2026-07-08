package repository

import (
	"tugas3/entity"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct{
	DB *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository *CategoryRepositoryImpl) Add(category entity.Categori) (entity.Categori, error){
	result := repository.DB.Create(&category)
	if result.Error != nil {
		return category, result.Error
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) Update(categori entity.Categori) (entity.Categori, error){
	result := repository.DB.Model(&entity.Categori{}).Where("id = ?", categori.ID).Updates(categori)
	if result.Error != nil {
		return categori, result.Error
	}
	return categori, nil
}

func (repository *CategoryRepositoryImpl) Delete(id int32) (error){
	var categori entity.Categori
	result := repository.DB.Delete(&categori, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *CategoryRepositoryImpl) GetAllData() ([]entity.Categori, error){
	var categoris []entity.Categori
	result := repository.DB.Find(&categoris)
	if result.Error != nil {
		return categoris, result.Error
	}
	return categoris, nil
}