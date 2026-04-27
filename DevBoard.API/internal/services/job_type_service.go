package services

import (
	"project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
)

type JobTypeService interface {
	CreateJobType(name string, userId uuid.UUID) (int, error)
	GetJobTypeById(id int) (*entities.JobType, error)
	GetAllJobTypes() ([]entities.JobType, error)
	UpdateJobType(id int, name string) error
	DeleteJobType(id int) error
}

type jobTypeService struct {
	repo repository.JobTypeRepository
}

func NewJobTypeService(repo repository.JobTypeRepository) JobTypeService {
	return &jobTypeService{repo: repo}
}

func (s *jobTypeService) CreateJobType(name string, userId uuid.UUID) (int, error) {
	jobType := &entities.JobType{
		Name: name,
		BaseEntity: entities.BaseEntity{
			CreatedBy:      userId,
			LastModifiedBy: userId,
		},
	}
	err := s.repo.Create(jobType)
	if err != nil {
		return 0, err
	}
	return jobType.Id, nil
}

func (s *jobTypeService) GetJobTypeById(id int) (*entities.JobType, error) {
	jobType, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if jobType == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return jobType, nil
}

func (s *jobTypeService) GetAllJobTypes() ([]entities.JobType, error) {
	return s.repo.ListAll()
}

func (s *jobTypeService) UpdateJobType(id int, name string) error {
	jobType, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if jobType == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	jobType.Name = name
	return s.repo.Update(jobType)
}

func (s *jobTypeService) DeleteJobType(id int) error {
	jobType, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if jobType == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return s.repo.Delete(jobType.Id)
}
