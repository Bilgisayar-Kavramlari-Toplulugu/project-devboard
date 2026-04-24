package repository

import (
	"project-devboard/internal/domain/entities"

	"gorm.io/gorm"
)

type SkillRepository interface {
	BaseRepository[entities.Skill, int]
}

type skillRepository struct {
	db *gorm.DB
	BaseRepository[entities.Skill, int]
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{
		BaseRepository: NewBaseRepository[entities.Skill, int](db),
		db:             db,
	}
}
