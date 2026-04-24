package services

import (
	"project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
)

type WorkLocationTypeService interface {
	CreateWorkLocationType(name string, userId uuid.UUID) (int, error)
	GetWorkLocationTypeById(id int) (*entities.WorkLocationType, error)
	GetAllWorkLocationTypes(limit, offset int) ([]entities.WorkLocationType, error)
	UpdateWorkLocationType(id int, name string) error
	DeleteWorkLocationType(id int) error
}

type workLocationTypeService struct {
	repo repository.WorkLocationTypeRepository
}

func NewWorkLocationTypeService(repo repository.WorkLocationTypeRepository) WorkLocationTypeService {
	return &workLocationTypeService{repo: repo}
}

func (s *workLocationTypeService) CreateWorkLocationType(name string, userId uuid.UUID) (int, error) {
	existing, err := s.repo.FindByName(name)
	if err != nil {
		return 0, err
	}
	if existing != nil {
		return 0, apperrors.New(apperrors.Conflict, apperrors.ErrAlreadyExists)
	}
	wlt := &entities.WorkLocationType{
		Name: name,
		BaseEntity: entities.BaseEntity{
			CreatedBy:      userId,
			LastModifiedBy: userId,
		},
	}
	if err := s.repo.Create(wlt); err != nil {
		return 0, err
	}
	return wlt.Id, nil
}

func (s *workLocationTypeService) GetWorkLocationTypeById(id int) (*entities.WorkLocationType, error) {
	wlt, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if wlt == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return wlt, nil
}

func (s *workLocationTypeService) GetAllWorkLocationTypes(limit, offset int) ([]entities.WorkLocationType, error) {
	return s.repo.List(limit, offset)
}

func (s *workLocationTypeService) UpdateWorkLocationType(id int, name string) error {
	wlt, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if wlt == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	existing, err := s.repo.FindByName(name)
	if err != nil {
		return err
	}
	if existing != nil && existing.Id != id {
		return apperrors.New(apperrors.Conflict, apperrors.ErrAlreadyExists)
	}
	wlt.Name = name
	return s.repo.Update(wlt)
}

func (s *workLocationTypeService) DeleteWorkLocationType(id int) error {
	wlt, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if wlt == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return s.repo.Delete(wlt.Id)
}
