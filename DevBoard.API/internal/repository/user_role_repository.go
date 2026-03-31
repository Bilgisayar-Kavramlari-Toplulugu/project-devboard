package repository

import (
	domain "project-devboard/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	BaseRepository[domain.UserRole, uuid.UUID]
	GetByUserID(userID uuid.UUID) ([]domain.UserRole, error)
	GetByRoleID(roleID uuid.UUID) ([]domain.UserRole, error)
	GetByUserAndRole(userID, roleID uuid.UUID) (*domain.UserRole, error)
}

type userRoleRepository struct {
	BaseRepository[domain.UserRole, uuid.UUID]
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{
		BaseRepository: NewBaseRepository[domain.UserRole, uuid.UUID](db),
		db:             db,
	}
}

func (r *userRoleRepository) GetByUserID(userID uuid.UUID) ([]domain.UserRole, error) {
	var userRoles []domain.UserRole
	err := r.db.Where("user_id = ?", userID).Find(&userRoles).Error
	return userRoles, err
}

func (r *userRoleRepository) GetByRoleID(roleID uuid.UUID) ([]domain.UserRole, error) {
	var userRoles []domain.UserRole
	err := r.db.Where("role_id = ?", roleID).Find(&userRoles).Error
	return userRoles, err
}

func (r *userRoleRepository) GetByUserAndRole(userID, roleID uuid.UUID) (*domain.UserRole, error) {
	var userRole domain.UserRole
	err := r.db.Where("user_id = ? AND role_id = ?", userID, roleID).First(&userRole).Error
	if err != nil {
		return nil, err
	}
	return &userRole, nil
}
