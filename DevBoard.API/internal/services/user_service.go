package services

import (
	"context"

	"project-devboard/internal/domain/entities"
	"project-devboard/internal/dtos"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

)

type UserService interface {
	CreateUser(ctx context.Context, req dtos.UserCreateRequest, actorID uuid.UUID) (*entities.User, error)
	GetUser(id uuid.UUID) (*entities.User, error)
	ListUsers(limit, offset int) ([]entities.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, req dtos.UserUpdateRequest, actorID uuid.UUID) (*entities.User, error)
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req dtos.UserCreateRequest, actorID uuid.UUID) (*entities.User, error) {
	existing, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if existing != nil {
		return nil, apperrors.New(apperrors.Conflict, apperrors.ErrUserAlreadyExists)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	userID := uuid.New()
	user := &entities.User{
		Id:               userID,
		Email:            req.Email,
		Firstname:        req.Firstname,
		Lastname:         req.Lastname,
		PhoneNumber:      req.PhoneNumber,
		Password:         string(hashed),
		IsEmailValidated: false,
		BaseEntity: entities.BaseEntity{
			IsActive:       true,
			CreatedBy:      actorID,
			LastModifiedBy: actorID,
		},
	}

	if err := s.repo.Create(user); err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	user.Password = ""
	return user, nil
}

func (s *userService) GetUser(id uuid.UUID) (*entities.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if user == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrUserNotFound)
	}
	return user, nil
}

func (s *userService) ListUsers(limit, offset int) ([]entities.User, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.List(limit)
}

func (s *userService) UpdateUser(ctx context.Context, id uuid.UUID, req dtos.UserUpdateRequest, actorID uuid.UUID) (*entities.User, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if existing == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrUserNotFound)
	}

	if req.Email != nil {
		emailOwner, err := s.repo.GetByEmail(*req.Email)
		if err != nil {
			return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
		}
		if emailOwner != nil && emailOwner.Id != id {
			return nil, apperrors.New(apperrors.Conflict, apperrors.ErrUserAlreadyExists)
		}
		existing.Email = *req.Email
	}

	if req.Firstname != nil {
		existing.Firstname = *req.Firstname
	}
	if req.Lastname != nil {
		existing.Lastname = *req.Lastname
	}
	if req.PhoneNumber != nil {
		existing.PhoneNumber = req.PhoneNumber
	}
	if req.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
		}
		existing.Password = string(hashed)
	}

	existing.BaseEntity.LastModifiedBy = actorID

	if err := s.repo.Update(existing); err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	existing.Password = ""
	return existing, nil
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if existing == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrUserNotFound)
	}

	return s.repo.Delete(id)
}
