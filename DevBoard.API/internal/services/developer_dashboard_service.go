package services

import (
	"project-devboard/internal/dtos"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
)

type DeveloperDashboardService interface {
	GetDashboardData(username string) (*dtos.DeveloperDashboardResponse, error)
	GetCurrentUserDashboardData(id uuid.UUID) (*dtos.CurrentUserDashboardResponse, error)
}

type developerDashboardService struct {
	repo repository.DeveloperDashboardRepository
}

func NewDeveloperDashboardService(repo repository.DeveloperDashboardRepository) DeveloperDashboardService {
	return &developerDashboardService{repo: repo}
}

func (s *developerDashboardService) GetDashboardData(username string) (*dtos.DeveloperDashboardResponse, error) {
	user, err := s.repo.GetUserProfile(username)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if user == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrUserNotFound)
	}
	res := dtos.NewDeveloperDashboardResponse(user)
	return &res, nil
}

func (s *developerDashboardService) GetCurrentUserDashboardData(id uuid.UUID) (*dtos.CurrentUserDashboardResponse, error) {
	user, err := s.repo.GetUserProfileByID(id)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if user == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrUserNotFound)
	}
	res := dtos.NewCurrentUserDashboardResponse(user)
	return &res, nil
}
