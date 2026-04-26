package repository

import (
	"project-devboard/internal/domain/entities"

	"gorm.io/gorm"
)

type CountryRepository interface {
	BaseRepository[entities.Country, int]
	ListAllAlphabetical() ([]entities.Country, error)
}
type countryRepository struct {
	BaseRepository[entities.Country, int]
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) CountryRepository {
	return &countryRepository{
		BaseRepository: NewBaseRepository[entities.Country, int](db),
		db:             db,
	}
}

func (r *countryRepository) ListAllAlphabetical() ([]entities.Country, error) {
	var countries []entities.Country
	err := r.db.Order("name asc").Find(&countries).Error
	return countries, err
}
