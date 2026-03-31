package repository

import (
	domain "project-devboard/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	BaseRepository[domain.Role, uuid.UUID]
	GetByName(name string) (*domain.Role, error)
}

type roleRepository struct {
	BaseRepository[domain.Role, uuid.UUID]
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		BaseRepository: NewBaseRepository[domain.Role, uuid.UUID](db),
		db:             db,
	}
}

func (r *roleRepository) GetByName(name string) (*domain.Role, error) {
	var role domain.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
