package repository

import (
	domain "project-devboard/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeveloperDashboardRepository interface {
	GetUserProfile(id uuid.UUID) (*domain.User, error)
}

type developerDashboardRepository struct {
	db *gorm.DB
}

func NewDeveloperDashboardRepository(db *gorm.DB) DeveloperDashboardRepository {
	return &developerDashboardRepository{db: db}
}

func (r *developerDashboardRepository) GetUserProfile(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := r.db.
		Preload("City").
		Preload("UserSkills.Skill").
		Preload("Certificates").
		Preload("Experiences").
		Preload("Educations").
		Preload("ProfessionalProfiles").
		Preload("UserJobTypes.JobType").
		Preload("UserWorkLocationTypes.WorkLocationType").
		Preload("Projects").
		Preload("SentMessages").
		Preload("ReceivedMessages").
		Preload("PublicEndorsementsSent").
		Preload("ProjectEndorsementsSent").
		Preload("References").
		Where("id = ?", id).
		First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
