package repository

import (
	"errors"

	domain "project-devboard/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[domain.User, uuid.UUID]
	GetByIdentifier(identifier string) (*domain.User, error)
	GetByIdentifierWithRoles(identifier string) (*domain.User, error)
	WithTx(tx *gorm.DB) UserRepository
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

func (r *userRepository) WithTx(tx *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[domain.User, uuid.UUID](tx),
		db:             tx,
	}
}

func (r *userRepository) GetByIdentifier(identifier string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByIdentifierWithRoles(identifier string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("UserRoles.Role").Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}