package services

import (
	"project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
)

type DeveloperDashboardService interface {
	GetDashboardData(userID uuid.UUID) (*entities.User, error)
}

type developerDashboardService struct {
	repo repository.DeveloperDashboardRepository
}

func NewDeveloperDashboardService(repo repository.DeveloperDashboardRepository) DeveloperDashboardService {
	return &developerDashboardService{repo: repo}
}

func (s *developerDashboardService) GetDashboardData(userID uuid.UUID) (*entities.User, error) {
	user, err := s.repo.GetUserProfile(userID)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if user == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrUserNotFound)
	}
	return user, nil
}
