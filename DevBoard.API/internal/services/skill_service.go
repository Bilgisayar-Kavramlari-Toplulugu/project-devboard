package services

import (
	"project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
)

//Kabul edilip beğenilirse service için aldığımız nesnelerin sonu input olsun burada onu kullandım beğenilmesi durumunda bütün serviceler için implement edebiliriz

type CreateSkillInput struct {
	Name        string
	SkillTypeId int
	UserId      uuid.UUID
}

type UpdateSkillInput struct {
	Id          int
	Name        string
	SkillTypeId int
}

type SkillService interface {
	CreateSkill(input CreateSkillInput) (int, error)
	GetSkillById(id int) (*entities.Skill, error)
	GetAllSkills() ([]entities.Skill, error)
	UpdateSkill(input UpdateSkillInput) error
	DeleteSkill(id int) error
}

type skillService struct {
	repository repository.SkillRepository
	//TODO: SkillTypeRepository'e ihtiyaç var mı kontrol edilecek
	//SkillTypeRepository eklenmesinin sebebi skill oluşturulurken skill type id'sinin doğruluğunun kontrol edilmesi düşünülmüştü ancak bu kontrolün yapılması gerekip gerekmediği henüz net değil
	skillTypeRepository repository.SkillTypeRepository
}

func NewSkillService(repo repository.SkillRepository, skillTypeRepo repository.SkillTypeRepository) SkillService {
	return &skillService{
		repository:          repo,
		skillTypeRepository: skillTypeRepo,
	}
}

func (s *skillService) CreateSkill(input CreateSkillInput) (int, error) {
	skillType, err := s.skillTypeRepository.GetByID(input.SkillTypeId)
	if err != nil {
		return 0, err
	}
	if skillType == nil {
		//ErrBadRequest oluşturup kullandım
		return 0, apperrors.New(apperrors.BadRequest, apperrors.ErrBadRequest)
	}
	skill := &entities.Skill{
		Name:        input.Name,
		SkillTypeId: input.SkillTypeId,
		BaseEntity: entities.BaseEntity{
			CreatedBy: input.UserId,
		},
	}
	err = s.repository.Create(skill)
	if err != nil {
		return 0, err
	}
	return skill.Id, nil
}

func (s *skillService) GetSkillById(id int) (*entities.Skill, error) {
	skill, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	if skill == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return skill, nil
}

func (s *skillService) GetAllSkills() ([]entities.Skill, error) {
	return s.repository.List(1000, 0)
}

func (s *skillService) UpdateSkill(input UpdateSkillInput) error {
	skill, err := s.repository.GetByID(input.Id)
	if err != nil {
		return err
	}
	if skill == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	skill.Name = input.Name
	skill.SkillTypeId = input.SkillTypeId
	return s.repository.Update(skill)
}

func (s *skillService) DeleteSkill(id int) error {
	skill, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}
	if skill == nil {
		return apperrors.New(apperrors.NotFound, apperrors.ErrNotFound)
	}
	return s.repository.Delete(skill.Id)
}
