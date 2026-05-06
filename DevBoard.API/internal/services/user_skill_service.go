package services

import (
	"project-devboard/internal/domain/entities"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
)

type UserSkillService interface {
	GetUserSkills(userID uuid.UUID) ([]entities.UserSkill, error)
	CreateUserSkills(userID uuid.UUID, skillIDs []int) ([]entities.UserSkill, error)
	UpdateUserSkills(userID uuid.UUID, skillIDs []int) ([]entities.UserSkill, error)
}

type userSkillService struct {
	repository repository.UserSkillRepository
}

func NewUserSkillService(repository repository.UserSkillRepository) UserSkillService {
	return &userSkillService{
		repository: repository,
	}
}

func (s *userSkillService) GetUserSkills(userID uuid.UUID) ([]entities.UserSkill, error) {
	userSkills, err := s.repository.GetByUserID(userID)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	return userSkills, nil
}

func (s *userSkillService) CreateUserSkills(userID uuid.UUID, skillIDs []int) ([]entities.UserSkill, error) {

	existing, err := s.repository.GetByUserID(userID)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	existingBySkill := make(map[int]struct{}, len(existing))
	for _, item := range existing {
		existingBySkill[item.Skill.Id] = struct{}{}
	}

	uniqueRequested := dedupeInts(skillIDs)

	if err := s.repository.ValidateIds(uniqueRequested); err != nil {
		return nil, apperrors.Wrap(apperrors.InvalidRequest, apperrors.ErrInvalidRequest, err)
	}

	toCreate := make([]entities.UserSkill, 0, len(uniqueRequested))
	for _, skillID := range uniqueRequested {
		if _, exists := existingBySkill[skillID]; exists {
			continue
		}

		toCreate = append(toCreate, entities.UserSkill{
			UserId:  userID,
			SkillId: skillID,
			BaseEntity: entities.BaseEntity{
				CreatedBy:      userID,
				LastModifiedBy: userID,
			},
		})
	}

	if err := s.repository.CreateBatch(toCreate); err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	return s.GetUserSkills(userID)
}

// Delete service fonksiyonu yazmama sebebim zaten update ile sadece gönderilen veriler kullanıcıya atanıyor geri kalanlar siliniyor.
func (s *userSkillService) UpdateUserSkills(userID uuid.UUID, skillIDs []int) ([]entities.UserSkill, error) {

	uniqueRequested := dedupeInts(skillIDs)

	if err := s.repository.ValidateIds(uniqueRequested); err != nil {
		return nil, apperrors.Wrap(apperrors.InvalidRequest, apperrors.ErrInvalidRequest, err)
	}

	if err := s.repository.DeleteByUserID(userID); err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	toCreate := make([]entities.UserSkill, 0, len(uniqueRequested))
	for _, skillID := range uniqueRequested {
		toCreate = append(toCreate, entities.UserSkill{
			UserId:  userID,
			SkillId: skillID,
			BaseEntity: entities.BaseEntity{
				CreatedBy:      userID,
				LastModifiedBy: userID,
			},
		})
	}

	if err := s.repository.CreateBatch(toCreate); err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	return s.GetUserSkills(userID)
}

// Birden aynı id değerini yok etmek için yardımcı fonksiyon. Normal şartlarda frontend'den böyle bir istek gelmemesi gerekir ancak yine de backend'de de kontrol etmek istedim.
func dedupeInts(values []int) []int {
	seen := make(map[int]struct{}, len(values))
	result := make([]int, 0, len(values))
	for _, value := range values {
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}
