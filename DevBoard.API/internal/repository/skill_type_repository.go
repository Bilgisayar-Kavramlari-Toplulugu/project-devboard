package repository

import (
	"project-devboard/internal/domain/entities"

	"gorm.io/gorm"
)

type SkillTypeRepository interface {
	BaseRepository[entities.SkillType, int]
}

type skillTypeRepository struct {
	BaseRepository[entities.SkillType, int]
	db *gorm.DB
}

func NewSkillTypeRepository(db *gorm.DB) SkillTypeRepository {
	return &skillTypeRepository{
		BaseRepository: NewBaseRepository[entities.SkillType, int](db),
		db:             db,
	}
}
