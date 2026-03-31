package repository

import (
	domain "project-devboard/internal/domain/entities"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[domain.User, uuid.UUID] // Inherit CRUD operations
	GetByEmail(email string) (*domain.User, error)
}

type userRepository struct {
	BaseRepository[domain.User, uuid.UUID]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[domain.User, uuid.UUID](db),
		db:             db,
	}
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
