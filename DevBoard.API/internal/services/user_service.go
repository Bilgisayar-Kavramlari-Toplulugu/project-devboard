package services

import (
	"context"

	"project-devboard/internal/domain/entities"
	"project-devboard/internal/dtos"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"
	"project-devboard/pkg/pagination"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, req dtos.UserCreateRequest, actorID uuid.UUID) (*entities.User, error)
	GetUser(id uuid.UUID) (*entities.User, error)
	ListUsers(page, pageSize int) (*pagination.PaginatedResult[dtos.UserResponse], error)
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
	existing, err := s.repo.GetByIdentifier(req.Email)
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

func (s *userService) ListUsers(page, pageSize int) (*pagination.PaginatedResult[dtos.UserResponse], error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	userlist, err := s.repo.PaginatedList(page, pageSize)
	userresponslist := &pagination.PaginatedResult[dtos.UserResponse]{
		Data:       dtos.NewUserResponses(userlist.Data),
		Total:      userlist.Total,
		Page:       userlist.Page,
		PageSize:   userlist.PageSize,
		TotalPages: userlist.TotalPages,
	}

	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	return userresponslist,nil
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
		emailOwner, err := s.repo.GetByIdentifier(*req.Email)
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
