package repository

import (
	"project-devboard/internal/domain/entities"

	"gorm.io/gorm"
)

type CityRepository interface {
	BaseRepository[entities.City, int]
	GetByCountryId(countryId int) ([]entities.City, error)
}

type cityRepository struct {
	BaseRepository[entities.City, int]
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityRepository{
		BaseRepository: NewBaseRepository[entities.City, int](db),
		db:             db,
	}
}

func (r *cityRepository) GetByCountryId(countryId int) ([]entities.City, error) {
	var cities []entities.City
	err := r.db.Where("country_id = ?", countryId).Order("name asc").Find(&cities).Error
	return cities, err
}