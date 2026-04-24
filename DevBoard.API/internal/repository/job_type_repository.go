package repository

import (
	"errors"

	"project-devboard/internal/domain/entities"

	"gorm.io/gorm"
)

type JobTypeRepository interface {
	BaseRepository[entities.JobType, int]
	FindByName(name string) (*entities.JobType, error)
}

type jobTypeRepository struct {
	BaseRepository[entities.JobType, int]
	db *gorm.DB
}

func NewJobTypeRepository(db *gorm.DB) JobTypeRepository {
	return &jobTypeRepository{
		BaseRepository: NewBaseRepository[entities.JobType, int](db),
		db:             db,
	}
}

func (r *jobTypeRepository) FindByName(name string) (*entities.JobType, error) {
	var jt entities.JobType
	err := r.db.Where("name = ?", name).First(&jt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// bulunamadı ama hata değil
		return nil, nil
	}
	if err != nil {
		// gerçek DB hatası
		return nil, err
	}
	return &jt, nil
}
