package services

import (
	"project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
)

type SkillTypeService interface {
	CreateSkillType(name string, userId uuid.UUID) (int, error)
	GetSkillTypeById(id int) (*entities.SkillType, error)
	GetAllSkillTypes() ([]entities.SkillType, error)
	UpdateSkillType(id int, name string) error
	DeleteSkillType(id int) error
}

type skillTypeService struct {
	repo repository.SkillTypeRepository
}

func NewSkillTypeService(repo repository.SkillTypeRepository) SkillTypeService {
	return &skillTypeService{repo: repo}
}

func (s *skillTypeService) CreateSkillType(name string, userId uuid.UUID) (int, error) {
	skillType := &entities.SkillType{
		Name: name,
		BaseEntity: entities.BaseEntity{
			CreatedBy: userId,
		},
	}
	err := s.repo.Create(skillType)
	if err != nil {
		return 0, err
	}
	return skillType.Id, nil
}

func (s *skillTypeService) GetSkillTypeById(id int) (*entities.SkillType, error) {
	skillType, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if skillType == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return skillType, nil
}

func (s *skillTypeService) GetAllSkillTypes() ([]entities.SkillType, error) {
	return s.repo.ListAll()
}

func (s *skillTypeService) UpdateSkillType(id int, name string) error {
	skillType, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if skillType == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	skillType.Name = name
	return s.repo.Update(skillType)
}

func (s *skillTypeService) DeleteSkillType(id int) error {
	skillType, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if skillType == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return s.repo.Delete(skillType.Id)
}
