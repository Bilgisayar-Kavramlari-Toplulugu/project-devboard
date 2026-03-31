package services

import (
	"errors"

	domain "project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"

	"github.com/google/uuid"
)

type UserRoleService interface {
	CreateUserRole(userRole *domain.UserRole) error
	GetUserRole(id uuid.UUID) (*domain.UserRole, error)
	ListUserRoles(limit, offset int) ([]domain.UserRole, error)
	UpdateUserRole(userRole *domain.UserRole) error
	DeleteUserRole(id uuid.UUID) error
	GetByUserID(userID uuid.UUID) ([]domain.UserRole, error)
	GetByRoleID(roleID uuid.UUID) ([]domain.UserRole, error)
}

type userRoleService struct {
	repo repository.UserRoleRepository
}

func NewUserRoleService(repo repository.UserRoleRepository) UserRoleService {
	return &userRoleService{repo: repo}
}

func (s *userRoleService) CreateUserRole(userRole *domain.UserRole) error {
	// Check if user-role assignment already exists
	existing, _ := s.repo.GetByUserAndRole(userRole.UserID, userRole.RoleID)
	if existing != nil {
		return errors.New("user role assignment already exists")
	}
	return s.repo.Create(userRole)
}

func (s *userRoleService) GetUserRole(id uuid.UUID) (*domain.UserRole, error) {
	return s.repo.GetByID(id)
}

func (s *userRoleService) ListUserRoles(limit, offset int) ([]domain.UserRole, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.List(limit, offset)
}

func (s *userRoleService) UpdateUserRole(userRole *domain.UserRole) error {
	existing, err := s.repo.GetByID(userRole.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("user role not found")
	}
	return s.repo.Update(userRole)
}

func (s *userRoleService) DeleteUserRole(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *userRoleService) GetByUserID(userID uuid.UUID) ([]domain.UserRole, error) {
	return s.repo.GetByUserID(userID)
}

func (s *userRoleService) GetByRoleID(roleID uuid.UUID) ([]domain.UserRole, error) {
	return s.repo.GetByRoleID(roleID)
}
