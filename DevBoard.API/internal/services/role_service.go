package services

import (
	"errors"

	domain "project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"

	"github.com/google/uuid"
)

type RoleService interface {
	CreateRole(role *domain.Role) error
	GetRole(id uuid.UUID) (*domain.Role, error)
	GetByName(name string) (*domain.Role, error)
	ListRoles(limit, offset int) ([]domain.Role, error)
	UpdateRole(role *domain.Role) error
	DeleteRole(id uuid.UUID) error
}

type roleService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) CreateRole(role *domain.Role) error {
	// Business logic: Check if role name already exists
	existing, _ := s.repo.GetByName(role.Name)
	if existing != nil {
		return errors.New("role name already exists")
	}
	return s.repo.Create(role)
}

func (s *roleService) GetRole(id uuid.UUID) (*domain.Role, error) {
	return s.repo.GetByID(id)
}

func (s *roleService) GetByName(name string) (*domain.Role, error) {
	return s.repo.GetByName(name)
}

func (s *roleService) ListRoles(limit, offset int) ([]domain.Role, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.List(limit, offset)
}

func (s *roleService) UpdateRole(role *domain.Role) error {
	existing, err := s.repo.GetByID(role.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("role not found")
	}
	return s.repo.Update(role)
}

func (s *roleService) DeleteRole(id uuid.UUID) error {
	return s.repo.Delete(id)
}
