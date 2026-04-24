package repository

import (
	"errors"

	"project-devboard/internal/domain/entities"

	"gorm.io/gorm"
)

type WorkLocationTypeRepository interface {
	BaseRepository[entities.WorkLocationType, int]
	FindByName(name string) (*entities.WorkLocationType, error)
}

type workLocationTypeRepository struct {
	BaseRepository[entities.WorkLocationType, int]
	db *gorm.DB
}

func NewWorkLocationTypeRepository(db *gorm.DB) WorkLocationTypeRepository {
	return &workLocationTypeRepository{
		BaseRepository: NewBaseRepository[entities.WorkLocationType, int](db),
		db:             db,
	}
}

func (r *workLocationTypeRepository) FindByName(name string) (*entities.WorkLocationType, error) {
	var wlt entities.WorkLocationType
	err := r.db.Where("name = ?", name).First(&wlt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &wlt, nil
}
