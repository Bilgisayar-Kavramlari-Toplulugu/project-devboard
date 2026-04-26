package repository

import (
	"project-devboard/internal/domain/entities"

	"gorm.io/gorm"
)

type JobTypeRepository interface {
	BaseRepository[entities.JobType, int]
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
