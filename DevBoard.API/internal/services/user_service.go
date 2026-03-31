package services

import (
	"errors"

	"project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *entities.User) error
	GetUser(id uuid.UUID) (*entities.User, error)
	ListUsers(limit, offset int) ([]entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *entities.User) error {
	// Business logic
	existing, _ := s.repo.GetByEmail(user.Email)
	if existing != nil {
		return errors.New("email already exists")
	}
	return s.repo.Create(user)
}

func (s *userService) GetUser(id uuid.UUID) (*entities.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) ListUsers(limit, offset int) ([]entities.User, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.List(limit, offset)
}

func (s *userService) UpdateUser(user *entities.User) error {
	existing, err := s.repo.GetByID(user.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("user not found")
	}
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.Delete(id)
}
